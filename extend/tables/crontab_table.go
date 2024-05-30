//go:build linux

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &CrontabTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type CrontabTable struct{}

func (p *CrontabTable) Name() string {
	return "crontab"
}

func (p *CrontabTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "event", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "minute", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "hour", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "day_of_month", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "month", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "day_of_week", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "command", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "path", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		//{Name: "pid_with_namespace", Type: table.INTEGER_TYPE, Options: table.ADDITIONAL | table.HIDDEN},
	}
}

func (p *CrontabTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenCrontab(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
