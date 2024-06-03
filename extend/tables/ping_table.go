package tables

import (
	"github.com/dean2021/sysql/extend/tables/networking"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &PingTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type PingTable struct{}

func (p *PingTable) Name() string {
	return "ping"
}

func (p *PingTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "addr", Type: table.TEXT_TYPE, Options: table.REQUIRED | table.INDEX},
		{Name: "count", Type: table.INTEGER_TYPE, Options: table.REQUIRED | table.INDEX},
		{Name: "ip", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "bytes", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "icmp_seq", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "time", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
	}
}

func (p *PingTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	rows, err := networking.GenPing(context)
	return rows, err
}
