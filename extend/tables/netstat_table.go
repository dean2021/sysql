package tables

import (
	"github.com/dean2021/sysql/extend/tables/networking"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &NetstatTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type NetstatTable struct{}

func (p *NetstatTable) Name() string {
	return "netstat"
}

func (p *NetstatTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "pid", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "local_port", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "remote_port", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "local_address", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "remote_address", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "protocol", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "family", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "fd", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "status", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *NetstatTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := networking.GenNetstat(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
