package system

import (
	"fmt"
	"github.com/dean2021/sysql/table"
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

func GenCPUInfo(context *table.QueryContext) (table.TableRows, error) {

	rows := table.TableRows{}

	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}

	for _, i := range info {
		rows = append(rows, table.TableRow{
			"device_id":           fmt.Sprintf("%v", i.CPU),
			"model":               i.Model,
			"manufacturer":        i.VendorID,
			"number_of_cores":     i.Cores,
			"current_clock_speed": i.Mhz,
			"load_percentage":     cpuPercent[i.CPU],
		})
	}

	return rows, nil
}
