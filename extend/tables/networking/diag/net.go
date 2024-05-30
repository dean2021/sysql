package diag

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/dean2021/sysql/extend/tables/common"

	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"syscall"
)

// Addr is implemented compatibility to psutil
type Addr struct {
	IP   string `json:"ip"`
	Port uint32 `json:"port"`
}

type NetConnectionKindType struct {
	family   uint32
	sockType uint32
	filename string
}

type conn struct {
	Family   uint32
	SockType uint32
	Laddr    Addr
	Raddr    Addr
	Status   string
	Inode    uint32
}

// http://students.mimuw.edu.pl/lxr/source/include/net/tcp_states.h
var TCPStatuses = map[string]string{
	"01": "ESTABLISHED",
	"02": "SYN_SENT",
	"03": "SYN_RECV",
	"04": "FIN_WAIT1",
	"05": "FIN_WAIT2",
	"06": "TIME_WAIT",
	"07": "CLOSE",
	"08": "CLOSE_WAIT",
	"09": "LAST_ACK",
	"0A": "LISTEN",
	"0B": "CLOSING",
}

var KindTCP4 = NetConnectionKindType{
	family:   syscall.AF_INET,
	sockType: syscall.SOCK_STREAM,
	filename: "tcp",
}
var kindTCP6 = NetConnectionKindType{
	family:   syscall.AF_INET6,
	sockType: syscall.SOCK_STREAM,
	filename: "tcp6",
}
var KindUDP4 = NetConnectionKindType{
	family:   syscall.AF_INET,
	sockType: syscall.SOCK_DGRAM,
	filename: "udp",
}
var KindUDP6 = NetConnectionKindType{
	family:   syscall.AF_INET6,
	sockType: syscall.SOCK_DGRAM,
	filename: "udp6",
}
var KindUNIX = NetConnectionKindType{
	family:   syscall.AF_UNIX,
	filename: "unix",
}

var netConnectionKindMap = map[string][]NetConnectionKindType{
	"all":   {KindTCP4, kindTCP6, KindUDP4, KindUDP6, KindUNIX},
	"tcp":   {KindTCP4, kindTCP6},
	"tcp4":  {KindTCP4},
	"tcp6":  {kindTCP6},
	"udp":   {KindUDP4, KindUDP6},
	"udp4":  {KindUDP4},
	"udp6":  {KindUDP6},
	"unix":  {KindUNIX},
	"inet":  {KindTCP4, kindTCP6, KindUDP4, KindUDP6},
	"inet4": {KindTCP4, KindUDP4},
	"inet6": {kindTCP6, KindUDP6},
}

// 解析网络文件
func ParserInet(file string, kind NetConnectionKindType) ([]conn, error) {

	if strings.HasSuffix(file, "6") && !common.PathExists(file) {
		// IPv6 not supported, return empty.
		return []conn{}, nil
	}

	// Read the contents of the /proc file with a single read sys call.
	// This minimizes duplicates in the returned connections
	// For more info:
	// https://github.com/shirou/gopsutil/pull/361
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := bytes.Split(contents, []byte("\n"))

	var ret []conn
	// skip first line
	for _, line := range lines[1:] {
		l := strings.Fields(string(line))
		if len(l) < 10 {
			continue
		}
		laddr := l[1]
		raddr := l[2]
		status := l[3]
		inode := l[9]
		if kind.sockType == syscall.SOCK_STREAM {
			status = TCPStatuses[status]
		} else {
			status = "NONE"
		}
		la, err := decodeAddress(kind.family, laddr)
		if err != nil {
			continue
		}
		ra, err := decodeAddress(kind.family, raddr)
		if err != nil {
			continue
		}
		i, err := strconv.Atoi(inode)
		if err != nil {
			continue
		}
		ret = append(ret, conn{
			Family:   kind.family,
			SockType: kind.sockType,
			Laddr:    la,
			Raddr:    ra,
			Status:   status,
			Inode:    uint32(i),
		})
	}
	return ret, nil
}

// decodeAddress decode addresse represents addr in proc/net/*
// ex:
// "0500000A:0016" -> "10.0.0.5", 22
// "0085002452100113070057A13F025401:0035" -> "2400:8500:1301:1052:a157:7:154:23f", 53
func decodeAddress(family uint32, src string) (Addr, error) {
	t := strings.Split(src, ":")
	if len(t) != 2 {
		return Addr{}, fmt.Errorf("does not contain port, %s", src)
	}
	addr := t[0]
	port, err := strconv.ParseInt("0x"+t[1], 0, 64)
	if err != nil {
		return Addr{}, fmt.Errorf("invalid port, %s", src)
	}
	decoded, err := hex.DecodeString(addr)
	if err != nil {
		return Addr{}, fmt.Errorf("decode error, %s", err)
	}
	var ip net.IP
	// Assumes this is little_endian
	if family == syscall.AF_INET {
		ip = Reverse(decoded)
	} else { // IPv6
		ip, err = parseIPv6HexString(decoded)
		if err != nil {
			return Addr{}, err
		}
	}
	return Addr{
		IP:   ip.String(),
		Port: uint32(port),
	}, nil
}

// Reverse reverses array of bytes.
func Reverse(s []byte) []byte {
	return ReverseWithContext(context.Background(), s)
}

func ReverseWithContext(ctx context.Context, s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// parseIPv6HexString parse array of bytes to IPv6 string
func parseIPv6HexString(src []byte) (net.IP, error) {
	if len(src) != 16 {
		return nil, fmt.Errorf("invalid IPv6 string")
	}

	buf := make([]byte, 0, 16)
	for i := 0; i < len(src); i += 4 {
		r := Reverse(src[i : i+4])
		buf = append(buf, r...)
	}
	return net.IP(buf), nil
}
