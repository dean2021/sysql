//go:build windows

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &ServicesTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type ServicesTable struct{}

func (p *ServicesTable) Name() string {
	return "services"
}

func (p *ServicesTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "service_type", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "display_name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "status", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "pid", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "start_type", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "win32_exit_code", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
	}
}

func (p *ServicesTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenWindowsServices(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
