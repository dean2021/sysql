//go:build linux || darwin

package system

import (
	"github.com/dean2021/sysql/table"
	"github.com/shirou/gopsutil/v3/load"
	"runtime"
)

func GenLoads(context *table.QueryContext) (table.TableRows, error) {
	avg, err := load.Avg()
	if err != nil {
		return nil, err
	}
	cores := runtime.NumCPU()
	return table.TableRows{
		table.TableRow{
			"period":  "1",
			"average": avg.Load1,
			"cores":   cores,
		},
		table.TableRow{
			"period":  "5",
			"average": avg.Load5,
			"cores":   cores,
		},
		table.TableRow{
			"period":  "15",
			"average": avg.Load15,
			"cores":   cores,
		},
	}, nil
}
