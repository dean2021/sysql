//go:build linux

package system

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/dean2021/sysql/table"
)

var authorizedKeyFileNames = []string{"authorized_keys", "authorized_keys2"}

func GenSSHKeys(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	users, err := getUsers()
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		for _, name := range authorizedKeyFileNames {
			//判断是否存在.ssh文件
			path := filepath.Join(u.Directory, ".ssh", name)
			stat, err := os.Stat(path)
			if err != nil {
				continue
			}
			keys, err := getPublicKeys(path)
			if err != nil {
				return nil, err
			}
			for _, key := range keys {
				results = append(results, table.TableRow{
					"uid":       u.Uid,
					"path":      path,
					"username":  u.Username,
					"file_name": name,
					"file_size": stat.Size(),
					"mod_time":  stat.ModTime().Format("2006-01-02 15:04:05"),
					"key":       key,
				})
			}
		}
	}
	return results, nil
}

func getPublicKeys(path string) ([]string, error) {
	//遍历authorized_keys文件
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	var lines []string
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
