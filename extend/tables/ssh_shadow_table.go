//go:build linux

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &ShadowTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type ShadowTable struct{}

func (p *ShadowTable) Name() string {
	return "shadow"
}

func (p *ShadowTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "password_status", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "hash", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "last_change", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "min", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "max", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "warning", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "inactive", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "expire", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "flag", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
	}
}

func (p *ShadowTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenShadow(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
