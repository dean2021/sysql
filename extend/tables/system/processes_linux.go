package system

import (
	"bytes"
	"github.com/dean2021/sysql/extend/tables/common"
	os2 "github.com/dean2021/sysql/misc/os"
	"github.com/dean2021/sysql/table"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/tklauser/go-sysconf"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var kMSIn1CLKTCK int64

func init() {
	scClkTck, err := sysconf.Sysconf(sysconf.SC_CLK_TCK)
	if err != nil {
		return
	}
	kMSIn1CLKTCK = 1000 / scClkTck
}

// 第一列数字(RUID):实际用户ID,指的是进程执行者是谁.
// 第二列数字(EUID):有效用户ID,指进程执行时对文件的访问权限.
// 第三列数字(SUID):保存设置用户ID,作为effective user ID的副本,在执行exec调用时后能重新恢复原来的effectiv user ID.
// 第四列数字(FSUID):目前进程的文件系统的用户识别码.一般情况下,文件系统的用户识别码(fsuid)与有效的用户识别码(euid)是相同的.
func getUid(p *process.Process) (uid int32, euid int32, suid int32, fsuid int32, err error) {
	uids, err := p.Uids()
	if err != nil {
		return -1, -1, -1, -1, err
	}
	for i, id := range uids {
		switch i {
		case 0:
			uid = id
		case 1:
			euid = id
		case 2:
			suid = id
		case 3:
			fsuid = id
		}
	}
	return
}

func getGid(p *process.Process) (gid int32, egid int32, sgid int32, fsgid int32, err error) {
	gids, err := p.Gids()
	if err != nil {
		return -1, -1, -1, -1, err
	}
	for i, id := range gids {
		switch i {
		case 0:
			gid = id
		case 1:
			egid = id
		case 2:
			sgid = id
		case 3:
			fsgid = id
		}
	}
	return
}

func splitProcStat(content []byte) []string {
	nameStart := bytes.IndexByte(content, '(')
	nameEnd := bytes.LastIndexByte(content, ')')
	restFields := strings.Fields(string(content[nameEnd+2:])) // +2 skip ') '
	name := content[nameStart+1 : nameEnd]
	pid := strings.TrimSpace(string(content[:nameStart]))
	fields := make([]string, 3, len(restFields)+3)
	fields[1] = string(pid)
	fields[2] = string(name)
	fields = append(fields, restFields...)
	return fields
}

type ProcStat struct {
	Terminal   uint64
	Parent     int64
	UTime      float64
	STime      float64
	PGroup     int64
	State      string
	UserTime   int64
	SystemTime int64
	Nice       int64
	Threads    int64
	StartTime  int64
}

func GetProcStat(pid int32) (*ProcStat, error) {
	var statPath string

	statPath = common.HostProc(strconv.Itoa(int(pid)), "stat")

	contents, err := ioutil.ReadFile(statPath)
	if err != nil {
		return nil, err
	}
	// Indexing from one, as described in `man proc` about the file /proc/[pid]/stat
	fields := splitProcStat(contents)

	terminal, err := strconv.ParseUint(fields[7], 10, 64)
	if err != nil {
		return nil, err
	}

	ppid, err := strconv.ParseInt(fields[4], 10, 32)
	if err != nil {
		return nil, err
	}
	utime, err := strconv.ParseFloat(fields[14], 64)
	if err != nil {
		return nil, err
	}

	stime, err := strconv.ParseFloat(fields[15], 64)
	if err != nil {
		return nil, err
	}

	group, err := strconv.ParseInt(fields[5], 10, 32)
	if err != nil {
		return nil, err
	}

	state := fields[3]

	userTime, err := strconv.ParseInt(fields[14], 10, 32)
	if err != nil {
		return nil, err
	}

	systemTime, err := strconv.ParseInt(fields[15], 10, 64)
	if err != nil {
		return nil, err
	}

	nice, err := strconv.ParseInt(fields[19], 10, 32)
	if err != nil {
		return nil, err
	}

	threads, err := strconv.ParseInt(fields[20], 10, 32)
	if err != nil {
		return nil, err
	}

	startTime, err := strconv.ParseInt(fields[22], 10, 64)
	if err != nil {
		return nil, err
	}

	proc := &ProcStat{
		Terminal:   terminal,
		Parent:     ppid,
		UTime:      utime,
		STime:      stime,
		UserTime:   userTime,
		State:      state,
		PGroup:     group,
		SystemTime: systemTime,
		StartTime:  startTime,
		Threads:    threads,
		Nice:       nice,
	}
	return proc, nil
}

type ProcIO struct {
	ReadBytes           int64
	WriteBytes          int64
	CancelledWriteBytes int64
}

func GetProcIO(pid int32) (*ProcIO, error) {
	var ioPath string
	ioPath = common.HostProc(strconv.Itoa(int(pid)), "io")
	contents, err := ioutil.ReadFile(ioPath)
	if err != nil {
		return nil, err
	}
	pIO := &ProcIO{}
	for _, line := range strings.Split(string(contents), "\n") {

		detail := strings.Split(line, ":")
		if len(detail) != 2 {
			continue
		}

		if detail[0] == "read_bytes" {
			if pIO.ReadBytes, err = strconv.ParseInt(strings.TrimSpace(detail[1]), 10, 64); err != nil {
				return nil, err
			}
		}

		if detail[0] == "write_bytes" {
			if pIO.WriteBytes, err = strconv.ParseInt(strings.TrimSpace(detail[1]), 10, 64); err != nil {
				return nil, err
			}
		}

		if detail[0] == "cancelled_write_bytes" {
			if pIO.CancelledWriteBytes, err = strconv.ParseInt(strings.TrimSpace(detail[1]), 10, 64); err != nil {
				return nil, err
			}
		}
	}

	return pIO, nil
}

func genProcess(pid int32) table.TableRow {

	var (
		residentSize     = -1
		totalSize        = -1
		diskBytesRead    = -1
		diskBytesWritten = -1
	)

	stat, err := GetProcStat(pid)
	if err != nil {
		return nil
	}

	p, err := process.NewProcess(pid)
	if err != nil {
		log.Println(err)
		return nil
	}

	path, _ := p.Exe()
	name, err := p.Name()
	if err != nil {
		if path != "" {
			name = filepath.Base(path)
		}
	}

	cwd, _ := p.Cwd()
	cmdline, _ := p.Cmdline()
	startTime, _ := p.CreateTime()
	uid, eUid, sUid, _, _ := getUid(p)
	gid, eGid, sGid, _, _ := getGid(p)
	userTime := stat.UserTime * kMSIn1CLKTCK
	systemTime := stat.SystemTime * kMSIn1CLKTCK
	root, _ := os.Readlink(common.HostProc(strconv.Itoa(int(pid)), "root"))

	onDisk := 0
	exists, err := os2.PathExists(path)
	if err != nil {
		onDisk = -1
	} else {
		if exists {
			onDisk = 1
		}
	}

	info, err := p.MemoryInfo()
	if err == nil {
		residentSize = int(info.RSS)
		totalSize = int(info.VMS)
	}
	ioInfo, err := GetProcIO(pid)
	if err == nil {
		diskBytesRead = int(ioInfo.ReadBytes)
		diskBytesWritten = int(ioInfo.WriteBytes) - int(ioInfo.CancelledWriteBytes)
	}

	// Fields unique to HackQuery
	username, _ := p.Username()
	terminal, _ := p.Terminal()
	cpuPercent, _ := p.CPUPercent()

	row := table.TableRow{
		"pid":                p.Pid,
		"name":               name,
		"path":               path,
		"cmdline":            cmdline,
		"state":              stat.State,
		"cwd":                cwd,
		"root":               root,
		"uid":                uid,
		"gid":                gid,
		"euid":               eUid,
		"egid":               eGid,
		"suid":               sUid,
		"sgid":               sGid,
		"on_disk":            onDisk,
		"wired_size":         0, // No support for unpagable counters in linux.
		"resident_size":      residentSize,
		"total_size":         totalSize,
		"user_time":          userTime,
		"system_time":        systemTime,
		"disk_bytes_read":    diskBytesRead,
		"disk_bytes_written": diskBytesWritten,
		"start_time":         startTime / 1000,
		"parent":             stat.Parent,
		"pgroup":             stat.PGroup,
		"threads":            stat.Threads,
		"nice":               stat.Nice,
		"username":           username,
		"terminal":           terminal,
		"cpu_percent":        cpuPercent,
	}

	return row
}
