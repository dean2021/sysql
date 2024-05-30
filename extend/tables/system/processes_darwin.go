package system

import (
	"github.com/dean2021/sysql/extend/tables/common"
	"github.com/dean2021/sysql/table"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"path/filepath"
	"strconv"
)

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

func genProcess(pid int32) table.TableRow {
	p, err := process.NewProcess(pid)
	if err != nil {
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
	ppid, _ := p.Ppid()
	cmdline, _ := p.Cmdline()
	startTime, _ := p.CreateTime()
	uid, euid, suid, _, _ := getUid(p)
	gid, egid, sgid, _, _ := getGid(p)
	root, _ := os.Readlink(common.HostProc(strconv.Itoa(int(pid)), "root"))

	onDisk := 0
	exists := common.PathExists(path)
	if exists {
		onDisk = 1
	}

	row := table.TableRow{
		"pid":     p.Pid,
		"name":    name,
		"path":    path,
		"cmdline": cmdline,
		//"state":   status[0],
		"cwd":     cwd,
		"root":    root,
		"uid":     uid,
		"gid":     gid,
		"euid":    euid,
		"egid":    egid,
		"suid":    suid,
		"sgid":    sgid,
		"on_disk": onDisk,
		//{Name: "wired_size", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "resident_size", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "total_size", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "user_time", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "system_time", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "disk_bytes_read", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "disk_bytes_written", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		"start_time": startTime / 1000,
		"parent":     ppid,
		//{Name: "pgroup", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "threads", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		//{Name: "nice", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		//{Name: "elevated_token", Type: table.INTEGER_TYPE, Options: table.HIDDEN},
		//{Name: "secure_process", Type: table.INTEGER_TYPE, Options: table.HIDDEN},
		//{Name: "protection_type", Type: table.TEXT_TYPE, Options: table.HIDDEN},
		//{Name: "virtual_process", Type: table.INTEGER_TYPE, Options: table.HIDDEN},
		//{Name: "elapsed_time", Type: table.BIGINT_TYPE, Options: table.HIDDEN},
		//{Name: "handle_count", Type: table.BIGINT_TYPE, Options: table.HIDDEN},
		//{Name: "percent_processor_time", Type: table.BIGINT_TYPE, Options: table.HIDDEN},
		//{Name: "upid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "uppid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "cpu_type", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		//{Name: "cpu_subtype", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
	}

	return row
}
