//go:build windows

package system

import (
	"github.com/dean2021/sysql/table"
	"github.com/yusufpapurcu/wmi"
)

type CIM_Product struct {
	Name    string
	Vendor  string
	Version string
	Caption string
}

func GenWindowsProducts(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	var s []CIM_Product
	err := wmi.Query("SELECT * FROM CIM_Product", &s)
	if err != nil {
		return nil, err
	}
	for _, pack := range s {
		results = append(results, table.TableRow{
			"name":    pack.Name,
			"vendor":  pack.Vendor,
			"version": pack.Version,
			"caption": pack.Caption,
		})
	}
	return results, nil
}
