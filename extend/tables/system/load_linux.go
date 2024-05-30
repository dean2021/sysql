package system

import (
	"github.com/dean2021/sysql/table"
	"github.com/shirou/gopsutil/v3/load"
)

func GenLoads(context *table.QueryContext) (table.TableRows, error) {
	avg, err := load.Avg()
	if err != nil {
		return nil, err
	}
	return table.TableRows{
		table.TableRow{
			"period":  "1",
			"average": avg.Load1,
		},
		table.TableRow{
			"period":  "5",
			"average": avg.Load5,
		},
		table.TableRow{
			"period":  "15",
			"average": avg.Load15,
		},
	}, nil
}
