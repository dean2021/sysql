//go:build linux || darwin

package utility

import (
	"fmt"
	"github.com/dean2021/sysql/table"
	"os"
)

func LsPerms(mode int) string {
	rwx := []string{"0", "1", "2", "3", "4", "5", "6", "7"}
	var bits string
	bits += rwx[(mode>>9)&7]
	bits += rwx[(mode>>6)&7]
	bits += rwx[(mode>>3)&7]
	bits += rwx[(mode>>0)&7]
	return bits
}

func GenFile(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	paths := context.Constraints.GetAll("path", table.EQUALS)
	for _, path := range paths {
		f := fmt.Sprintf("%v", path)
		stat, err := os.Stat(f)
		if err != nil {
			continue
		}
		if stat.IsDir() {
			continue
		}

		result := genFile(f)
		if result != nil {
			results = append(results, result)
		}
	}

	directors := context.Constraints.GetAll("directory", table.EQUALS)
	for _, directory := range directors {
		dir := fmt.Sprintf("%v", directory)
		readDir, err := os.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		for _, fi := range readDir {
			if fi.IsDir() {
				continue
			}
			f := dir + string(os.PathSeparator) + fi.Name()
			result := genFile(f)
			if result != nil {
				results = append(results, result)
			}
		}
	}
	return results, nil
}
