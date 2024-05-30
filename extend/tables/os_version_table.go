package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &OsVersionTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type OsVersionTable struct{}

func (p *OsVersionTable) Name() string {
	return "os_version"
}

func (p *OsVersionTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "version", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "major", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "minor", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "patch", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "build", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "platform", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "platform_like", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "codename", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "arch", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "install_date", Type: table.BIGINT_TYPE, Options: table.HIDDEN},
		{Name: "pid_with_namespace", Type: table.INTEGER_TYPE, Options: table.ADDITIONAL | table.HIDDEN},
		{Name: "mount_namespace_id", Type: table.TEXT_TYPE, Options: table.HIDDEN},
	}
}

func (p *OsVersionTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	rows, err := system.GenOSVersion(context)
	return rows, err
}
