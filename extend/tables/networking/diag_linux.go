package networking

import (
	"bytes"
	"fmt"
	"github.com/dean2021/sysql/extend/tables/networking/diag"
	"github.com/dean2021/sysql/table"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func GenNetstatWithDiag(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	var (
		networks    []Network
		inodeMapper map[uint32]int
		err         error
	)

	inodeMapper, err = GetProcessInodeMapper(-1)
	if err != nil {
		return nil, err
	}

	// linux内核版本小于3.x.x不支持inet_diag_req_v2, 并且不支持获取udp协议的连接, 此处做了兼容
	majorVersion, err := GetLinuxKernelMajorVersion()
	if err != nil {
		return nil, err
	}
	if majorVersion < 3 {
		networks, err = GetNetworkInfoWithInetDiagReq(inodeMapper)
		if err != nil {
			return nil, err
		}
	} else {
		networks, err = GetNetworkInfoWithInetDiagReqV2(inodeMapper)
		if err != nil {
			return nil, err
		}
	}

	for _, connection := range networks {
		results = append(results, table.TableRow{
			"pid":            connection.Pid,
			"local_port":     connection.LocalPort,
			"remote_port":    connection.RemotePort,
			"local_address":  connection.LocalIP,
			"remote_address": connection.RemoteIP,
			"family":         connection.Family,
			"protocol":       connection.Proto,
			"inode":          connection.Inode,
			"status":         connection.Status,
		})
	}
	return results, nil
}

type Network struct {
	Proto      string `json:"proto"`
	LocalIP    string `json:"localIP"`
	LocalPort  string `json:"localPort"`
	RemoteIP   string `json:"remoteIP"`
	RemotePort string `json:"remotePort"`
	Status     string `json:"status"`
	Pid        int32  `json:"pid"`
	Family     uint8  `json:"family"`
	Inode      uint32 `json:"inode"`
}

// GetLinuxKernelVersion 获取linux内核版本
func GetLinuxKernelVersion() (string, error) {
	var utsName unix.Utsname
	err := unix.Uname(&utsName)
	if err != nil {
		return "", err
	}
	return string(utsName.Release[:bytes.IndexByte(utsName.Release[:], 0)]), nil
}

// GetLinuxKernelMajorVersion 获取linux内核版本主版本号
func GetLinuxKernelMajorVersion() (int, error) {
	version, err := GetLinuxKernelVersion()
	if err != nil {
		return 0, err
	}
	versions := strings.Split(version, ".")
	majorVersion, err := strconv.Atoi(versions[0])
	if err != nil {
		return 0, err
	}
	return majorVersion, nil
}

func GetProcessInodeMapper(max int) (map[uint32]int, error) {
	root := os.Getenv("PROC_ROOT")
	if root == "" {
		root = "/proc"
	}
	dir, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	var processInodes = make(map[uint32]int, 1)
	for _, d := range dir {
		if !d.IsDir() {
			continue
		}
		pid, err := strconv.Atoi(d.Name())
		if err != nil {
			continue
		}
		// skip self process
		if pid == os.Getpid() {
			continue
		}
		pidDir := filepath.Join(root, d.Name())
		fdPath := filepath.Join(pidDir, "fd")
		d, err := os.Open(fdPath)
		if err != nil {
			continue
		}
		defer d.Close()
		names, err := d.Readdirnames(max)
		if err != nil {
			continue
		}
		for _, fd := range names {
			link, err := os.Readlink(filepath.Join(fdPath, fd))
			if err != nil {
				continue
			}
			inode, err := parseSocketInode(link)
			if err != nil {
				continue
			}
			if inode == 0 {
				continue
			}
			processInodes[inode] = pid
		}
	}
	return processInodes, nil
}

func parseSocketInode(lnk string) (uint32, error) {
	const pattern = "socket:["
	ind := strings.Index(lnk, pattern)
	if ind == -1 {
		return 0, nil
	}
	var ino uint32
	n, err := fmt.Sscanf(lnk, "socket:[%d]", &ino)
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, fmt.Errorf("'%s' should be pattern '[socket:\\%d]'", lnk)
	}
	return ino, nil
}

// GetNetworkInfoWithInetDiagReqV2 基于InetDiagReqV2的获取网络信息
// 仅在linux内核 > 3.x.x可用
func GetNetworkInfoWithInetDiagReqV2(inodeMapper map[uint32]int) ([]Network, error) {
	var (
		pid      int32 = -1
		networks []Network
	)
	protocols := map[string]uint8{
		"tcp": syscall.IPPROTO_TCP,
		"udp": syscall.IPPROTO_UDP,
	}
	for protocolName, protocol := range protocols {
		req := diag.NewInetDiagReqV2(diag.AF_INET, protocol)
		messages, err := diag.NetlinkInetDiagWithBuf(req, nil, nil)
		if err != nil {
			return nil, err
		}
		for _, msg := range messages {
			if v, ok := inodeMapper[msg.Inode]; ok {
				pid = int32(v)
			} else {
				pid = -1
			}
			networks = append(networks, Network{
				Proto:      protocolName,
				LocalIP:    msg.SrcIP().String(),
				LocalPort:  strconv.Itoa(msg.SrcPort()),
				RemoteIP:   msg.DstIP().String(),
				RemotePort: strconv.Itoa(msg.DstPort()),
				Status:     diag.TCPState(msg.State).String(),
				Pid:        pid,
				Family:     msg.Family,
				Inode:      msg.Inode,
			})
		}
	}
	return networks, nil
}

// GetNetworkInfoWithInetDiagReq 基于InetDiagReq的获取网络信息
// 在linux内核2.x.x可用, UDP协议数据需要通过解析/proc/net/udp来实现
func GetNetworkInfoWithInetDiagReq(inodeMapper map[uint32]int) ([]Network, error) {
	var (
		pid      int32 = -1
		networks []Network
	)

	req := diag.NewInetDiagReq()
	messages, err := diag.NetlinkInetDiagWithBuf(req, nil, nil)
	if err != nil {
		return nil, err
	}
	for _, msg := range messages {
		if v, ok := inodeMapper[msg.Inode]; ok {
			pid = int32(v)
		} else {
			pid = -1
		}
		networks = append(networks, Network{
			Proto:      "tcp",
			LocalIP:    msg.SrcIP().String(),
			LocalPort:  strconv.Itoa(msg.SrcPort()),
			RemoteIP:   msg.DstIP().String(),
			RemotePort: strconv.Itoa(msg.DstPort()),
			Status:     diag.TCPState(msg.State).String(),
			Pid:        pid,
			Family:     msg.Family,
			Inode:      msg.Inode,
		})
	}

	// UDP协议数据需要通过解析/proc/net/udp来实现
	// udp连接一般较少， 故不存在cpu过高问题
	root := os.Getenv("PROC_ROOT")
	if root == "" {
		root = "/proc"
	}
	udpPath := filepath.Join(root, "net", "udp")
	connections, err := diag.ParserInet(udpPath, diag.KindUDP4)
	if err != nil {
		return networks, nil
	}
	for _, conn := range connections {
		if v, ok := inodeMapper[conn.Inode]; ok {
			pid = int32(v)
		} else {
			pid = -1
		}
		networks = append(networks, Network{
			Proto:      "udp",
			LocalIP:    conn.Laddr.IP,
			LocalPort:  strconv.Itoa(int(conn.Laddr.Port)),
			RemoteIP:   conn.Raddr.IP,
			RemotePort: strconv.Itoa(int(conn.Raddr.Port)),
			Status:     conn.Status,
			Pid:        pid,
			Family:     uint8(conn.Family),
			Inode:      conn.Inode,
		})
	}
	return networks, nil
}
