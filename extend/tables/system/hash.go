package system

import (
	"fmt"
	"github.com/dean2021/sysql/extend/functions"
	"github.com/dean2021/sysql/table"
	"os"
	"path/filepath"
)

func genHash(path string) table.TableRow {
	directory, _ := filepath.Abs(path)
	return table.TableRow{
		"path":      path,
		"directory": filepath.Dir(directory),
		"md5":       functions.Md5File(path),
		"sha1":      functions.Sha1(path),
		"sha256":    functions.Sha256(path),
	}
}

func GenHash(context *table.QueryContext) (table.TableRows, error) {
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
		results = append(results, genHash(f))
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
			results = append(results, genHash(f))
		}
	}
	return results, nil
}
