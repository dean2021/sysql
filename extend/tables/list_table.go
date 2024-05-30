package tables

import (
	"fmt"
	"github.com/dean2021/sysql/table"
	"os"
	"path/filepath"
)

func init() {
	t := &ListTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type ListTable struct{}

func (p *ListTable) Name() string {
	return "list"
}

func (p *ListTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "path", Type: table.TEXT_TYPE, Options: table.INDEX | table.REQUIRED},
		{Name: "directory", Type: table.TEXT_TYPE, Options: table.REQUIRED},
		{Name: "type", Type: table.TEXT_TYPE},
	}
}

func (p *ListTable) Generate(context *table.QueryContext) (table.TableRows, error) {

	var results table.TableRows

	directors := context.Constraints.GetAll("directory", table.EQUALS)
	for _, v := range directors {
		dir := fmt.Sprintf("%v", v)
		readDir, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		for _, fi := range readDir {
			path := dir + string(os.PathSeparator) + fi.Name()
			directory, _ := filepath.Abs(path)

			fType := "file"
			if fi.IsDir() {
				fType = "dir"
			}
			results = append(results, table.TableRow{
				"path":      path,
				"directory": filepath.Dir(directory),
				"type":      fType,
			})
		}
	}

	return results, nil
}
