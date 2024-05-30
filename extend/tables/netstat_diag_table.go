//go:build linux

package tables

import (
	"github.com/dean2021/sysql/extend/tables/networking"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &NetstatDiagTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type NetstatDiagTable struct{}

func (p *NetstatDiagTable) Name() string {
	return "netstat_diag"
}

func (p *NetstatDiagTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "pid", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "local_port", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "remote_port", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "local_address", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "remote_address", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "protocol", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "family", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "inode", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "status", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *NetstatDiagTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := networking.GenNetstatWithDiag(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
