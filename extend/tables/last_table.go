package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &LastTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type LastTable struct{}

func (p *LastTable) Name() string {
	return "last"
}

func (p *LastTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "username", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "tty", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "pid", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "type", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "type_name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "time", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "host", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *LastTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	rows, err := system.GenLast(context)
	return rows, err
}
