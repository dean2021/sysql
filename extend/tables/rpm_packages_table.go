//go:build linux

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &RpmPackagesTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type RpmPackagesTable struct{}

func (p *RpmPackagesTable) Name() string {
	return "rpm_packages"
}

func (p *RpmPackagesTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "version", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "release", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "description", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *RpmPackagesTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenRpmPackages(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
