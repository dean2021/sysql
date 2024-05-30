package networking

import (
	"github.com/dean2021/sysql/table"
	"github.com/shirou/gopsutil/v3/net"
)

func genNetstatImpl(context *table.QueryContext) (table.TableRows, error) {
	connections, err := net.Connections("all")
	if err != nil {
		return nil, err
	}
	var results table.TableRows
	for _, connection := range connections {
		results = append(results, table.TableRow{
			"pid":            connection.Pid,
			"local_port":     connection.Laddr.Port,
			"remote_port":    connection.Raddr.Port,
			"local_address":  connection.Laddr.IP,
			"remote_address": connection.Raddr.IP,
			"family":         connection.Family,
			"protocol":       connection.Type,
			"fd":             connection.Fd,
			"status":         connection.Status,
		})
	}
	return results, nil
}
