//go:build windows
// +build windows

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &LocalUserTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type LocalUserTable struct{}

func (p *LocalUserTable) Name() string {
	return "local_users"
}

func (p *LocalUserTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "name", Type: table.TEXT_TYPE, Options: table.INDEX},
		{Name: "description", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "enabled", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "principalSource", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "accountExpires", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "passwordChangeableDate", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "passwordExpires", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "userMayChangePassword", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "lastLogon", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
	}
}

func (p *LocalUserTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenLocalUsers(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
