//go:build windows
// +build windows

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &UsersTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type UsersTable struct{}

func (p *UsersTable) Name() string {
	return "users"
}

func (p *UsersTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "sid", Type: table.TEXT_TYPE, Options: table.INDEX},
		{Name: "name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "accountType", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "caption", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "description", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "disabled", Type: table.INTEGER_TYPE, Options: table.DEFAULT}, // Assuming boolean is represented as integer
		{Name: "domain", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "fullName", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "installDate", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "localAccount", Type: table.INTEGER_TYPE, Options: table.DEFAULT},       // Assuming boolean is represented as integer
		{Name: "lockout", Type: table.INTEGER_TYPE, Options: table.DEFAULT},            // Assuming boolean is represented as integer
		{Name: "passwordChangeable", Type: table.INTEGER_TYPE, Options: table.DEFAULT}, // Assuming boolean is represented as integer
		{Name: "passwordExpires", Type: table.INTEGER_TYPE, Options: table.DEFAULT},    // Assuming boolean is represented as integer
		{Name: "passwordRequired", Type: table.INTEGER_TYPE, Options: table.DEFAULT},   // Assuming boolean is represented as integer
		{Name: "sidType", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "status", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "passwordLastSet", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *UsersTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenUsers(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
