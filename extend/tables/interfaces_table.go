package tables

import (
	"errors"
	"github.com/dean2021/sysql/table"
	"net"
)

func init() {
	t := &InterfaceTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type InterfaceTable struct{}

func (p *InterfaceTable) Name() string {
	return "interfaces"
}

func (p *InterfaceTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "index", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "mtu", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "flags", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "addr", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "hardware_addr", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *InterfaceTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	var rows table.TableRows
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, errors.New("Failed to get interfaces: " + err.Error())
	}
	for _, inter := range interfaces {
		address, err := inter.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range address {
			rows = append(rows, map[string]interface{}{
				"name":          inter.Name,
				"index":         inter.Index,
				"mtu":           inter.MTU,
				"flags":         inter.Flags.String(),
				"addr":          addr.String(),
				"hardware_addr": inter.HardwareAddr.String(),
			})
		}
	}
	return rows, nil
}
