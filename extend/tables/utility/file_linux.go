package utility

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/dean2021/sysql/table"
	"golang.org/x/sys/unix"
)

func genFile(path string) table.TableRow {

	var stat unix.Stat_t
	err := unix.Stat(path, &stat)
	if err != nil {
		return nil
	}

	atime, _ := stat.Atim.Unix()
	mtime, _ := stat.Mtim.Unix()
	ctime, _ := stat.Ctim.Unix()

	lstat, err := os.Lstat(path)
	if err != nil {
		return nil
	}

	symlink := 0
	if lstat.Mode()&os.ModeSymlink == os.ModeSymlink {
		symlink = 1
	}

	directory, _ := filepath.Abs(path)

	chattr := checkChattr(path)

	return table.TableRow{
		"path":       path,
		"directory":  filepath.Dir(directory),
		"filename":   filepath.Base(path),
		"inode":      stat.Ino,
		"uid":        stat.Uid,
		"gid":        stat.Gid,
		"mode":       LsPerms(int(stat.Mode)),
		"device":     stat.Rdev,
		"size":       stat.Size,
		"block_size": stat.Blksize,
		"atime":      atime,
		"mtime":      mtime,
		"ctime":      ctime,
		"btime":      0,
		"hard_links": stat.Nlink,
		"symlink":    symlink,
		"chattr":     chattr,
		// TODO
		"type": "0",
	}
}

// CheckChattr 检查文件是否被chattr上锁
func checkChattr(filePath string) string {
	// 构造命令
	cmd := exec.Command("lsattr", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}
