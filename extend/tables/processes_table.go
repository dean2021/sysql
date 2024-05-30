package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &ProcessesTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type ProcessesTable struct{}

func (p *ProcessesTable) Name() string {
	return "processes"
}

func (p *ProcessesTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "pid", Type: table.BIGINT_TYPE, Options: table.INDEX},
		{Name: "name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "path", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "cmdline", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "state", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "cwd", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "root", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "uid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "gid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "euid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "egid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "suid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "sgid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "on_disk", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "wired_size", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "resident_size", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "total_size", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "user_time", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "system_time", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "disk_bytes_read", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "disk_bytes_written", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "start_time", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "parent", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "pgroup", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "threads", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "nice", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		//{Name: "elevated_token", Type: table.INTEGER_TYPE, Options: table.HIDDEN},
		//{Name: "secure_process", Type: table.INTEGER_TYPE, Options: table.HIDDEN},
		//{Name: "protection_type", Type: table.TEXT_TYPE, Options: table.HIDDEN},
		//{Name: "virtual_process", Type: table.INTEGER_TYPE, Options: table.HIDDEN},
		//{Name: "elapsed_time", Type: table.BIGINT_TYPE, Options: table.HIDDEN},
		//{Name: "handle_count", Type: table.BIGINT_TYPE, Options: table.HIDDEN},
		//{Name: "percent_processor_time", Type: table.BIGINT_TYPE, Options: table.HIDDEN},
		//{Name: "upid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "uppid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "cpu_type", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		//{Name: "cpu_subtype", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "username", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "terminal", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "cpu_percent", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
	}
}

func (p *ProcessesTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenProcesses(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
