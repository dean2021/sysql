//go:build linux

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
		{Name: "uid", Type: table.BIGINT_TYPE, Options: table.INDEX},
		{Name: "gid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "uid_signed", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		//{Name: "gid_signed", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "username", Type: table.TEXT_TYPE, Options: table.ADDITIONAL},
		{Name: "description", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "directory", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "shell", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		//{Name: "uuid", Type: table.TEXT_TYPE, Options: table.INDEX},
		//{Name: "type", Type: table.TEXT_TYPE, Options: table.HIDDEN},
		//{Name: "is_hidden", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		//{Name: "pid_with_namespace", Type: table.INTEGER_TYPE, Options: table.ADDITIONAL | table.HIDDEN},
	}
}

func (p *UsersTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenUsers(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
