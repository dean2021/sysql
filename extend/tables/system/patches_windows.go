package system

import (
	"github.com/dean2021/sysql/table"
	"github.com/yusufpapurcu/wmi"
)

type Win32QuickFix struct {
	CSName      string `json:"csname"`
	HotFixID    string `json:"hotfix_id"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
	FixComments string `json:"fix_comments"`
	InstalledBy string `json:"installed_by"`
	InstallDate string `json:"install_date"`
	InstalledOn string `json:"installed_on"`
}

func GenPatches(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	var s []Win32QuickFix
	err := wmi.Query("SELECT * FROM Win32_QuickFixEngineering", &s)
	if err != nil {
		return nil, err
	}
	for _, pack := range s {
		results = append(results, table.TableRow{
			"csname":       pack.CSName,
			"hotfix_id":    pack.HotFixID,
			"caption":      pack.Caption,
			"description":  pack.Description,
			"fix_comments": pack.FixComments,
			"installed_by": pack.InstalledBy,
			"install_date": pack.InstallDate,
			"installed_on": pack.InstalledOn,
		})
	}
	return results, nil
}
