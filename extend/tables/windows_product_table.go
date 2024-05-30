//go:build windows

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &WindowsProductTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type WindowsProductTable struct{}

func (p *WindowsProductTable) Name() string {
	return "windows_product"
}

func (p *WindowsProductTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "vendor", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "version", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "caption", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *WindowsProductTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenWindowsProducts(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
