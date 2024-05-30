package system

import (
	"github.com/dean2021/sysql/table"
	"github.com/shirou/gopsutil/v3/process"
)

func GenProcesses(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	var pidLists []int
	if context.Constraints.Count("pid") > 0 && context.Constraints.Exists("pid", table.EQUALS) {
		all := context.Constraints.GetAll("pid", table.EQUALS)
		for _, pid := range all {
			pidLists = append(pidLists, int(pid.(int64)))
		}
	} else {
		all, err := process.Pids()
		if err == nil {
			for _, pid := range all {
				pidLists = append(pidLists, int(pid))
			}
		}
	}
	for _, pid := range pidLists {
		row := genProcess(int32(pid))
		if row != nil {
			results = append(results, row)
		}
	}
	return results, nil
}
