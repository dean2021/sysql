//go:build linux || darwin

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &LoadAverageTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type LoadAverageTable struct{}

func (p *LoadAverageTable) Name() string {
	return "load_average"
}

func (p *LoadAverageTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "period", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "average", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "cores", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
	}
}

func (p *LoadAverageTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenLoads(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
