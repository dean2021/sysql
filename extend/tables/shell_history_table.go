//go:build linux

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &ShellHistoryTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type ShellHistoryTable struct{}

func (p *ShellHistoryTable) Name() string {
	return "shell_history"
}

func (p *ShellHistoryTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "uid", Type: table.BIGINT_TYPE, Options: table.INDEX},
		{Name: "time", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "command", Type: table.TEXT_TYPE, Options: table.ADDITIONAL},
		{Name: "history_file", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *ShellHistoryTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenShellHistory(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
