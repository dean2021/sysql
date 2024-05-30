package tables

import (
	"github.com/dean2021/sysql/extend/tables/networking"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &CurlTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type CurlTable struct{}

func (p *CurlTable) Name() string {
	return "curl"
}

func (p *CurlTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "url", Type: table.TEXT_TYPE, Options: table.REQUIRED | table.INDEX},
		{Name: "method", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "user_agent", Type: table.TEXT_TYPE, Options: table.ADDITIONAL},
		{Name: "response_code", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "round_trip_time", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "bytes", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "result", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *CurlTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	rows, err := networking.GenCurl(context)
	return rows, err
}
