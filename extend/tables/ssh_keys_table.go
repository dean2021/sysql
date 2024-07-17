//go:build linux

package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &SSHKeysTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type SSHKeysTable struct{}

func (p *SSHKeysTable) Name() string {
	return "ssh_keys"
}

func (p *SSHKeysTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "uid", Type: table.BIGINT_TYPE, Options: table.INDEX},
		{Name: "path", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "username", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "file_name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "file_size", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "mod_time", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "key", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *SSHKeysTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenSSHKeys(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
