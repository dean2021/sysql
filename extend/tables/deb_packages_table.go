//go:build linux

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &DebPackagesTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type DebPackagesTable struct{}

func (p *DebPackagesTable) Name() string {
	return "deb_packages"
}

func (p *DebPackagesTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "version", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "description", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "source", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "size", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "arch", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "revision", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "status", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "maintainer", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "section", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "priority", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "homepage", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "original_maintainer", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "replaces", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "provides", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "recommends", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "suggests", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "breaks", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		//{Name: "pid_with_namespace", Type: table.INTEGER_TYPE, Options: table.ADDITIONAL | table.HIDDEN},
	}
}

func (p *DebPackagesTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenDebPackages(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
