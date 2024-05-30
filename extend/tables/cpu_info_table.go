package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &CPUInfoTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type CPUInfoTable struct{}

func (p *CPUInfoTable) Name() string {
	return "cpu_info"
}

func (p *CPUInfoTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "device_id", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "model", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "manufacturer", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "number_of_cores", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "current_clock_speed", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "load_percentage", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
	}
}

func (p *CPUInfoTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenCPUInfo(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
