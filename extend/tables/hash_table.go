package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &HashTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type HashTable struct{}

func (p *HashTable) Name() string {
	return "hash"
}

func (p *HashTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "path", Type: table.TEXT_TYPE, Options: table.INDEX | table.REQUIRED},
		{Name: "directory", Type: table.TEXT_TYPE, Options: table.REQUIRED},
		{Name: "md5", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "sha1", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "sha256", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		//{Name: "pid_with_namespace", Type: table.INTEGER_TYPE, Options: table.ADDITIONAL | table.HIDDEN},
		//{Name: "mount_namespace_id", Type: table.TEXT_TYPE, Options: table.HIDDEN},
	}
}

func (p *HashTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	rows, err := system.GenHash(context)
	return rows, err
}
