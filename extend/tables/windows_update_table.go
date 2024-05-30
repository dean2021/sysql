//go:build windows

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &WindowsUpdateTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type WindowsUpdateTable struct{}

func (p *WindowsUpdateTable) Name() string {
	return "windows_update"
}

func (p *WindowsUpdateTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "title", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "description", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "categories", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "kb_article_ids", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "more_info_urls", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "support_url", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "update_id", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "revision_number", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "severity", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "last_deployment_change_time", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *WindowsUpdateTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenWindowsUpdates(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
