//go:build windows

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &PatchesTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type PatchesTable struct{}

func (p *PatchesTable) Name() string {
	return "patches"
}

func (p *PatchesTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "csname", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "hotfix_id", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "caption", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "description", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "fix_comments", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "installed_by", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "install_date", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "installed_on", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *PatchesTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenPatches(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
