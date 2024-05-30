package sysql

import (
	"database/sql"
	"database/sql/driver"
	"github.com/dean2021/go-sqlite3"
	"github.com/dean2021/sysql/extend/functions"
	_ "github.com/dean2021/sysql/extend/tables"
	"github.com/dean2021/sysql/table"
)

const DriverName = "SQLITE3_SYSQL_EXTENSIONS"

func Initialize() {
	sql.Register(DriverName, &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			for name, function := range functions.Functions {
				err := conn.RegisterFunc(name, function, true)
				if err != nil {
					return err
				}
			}
			tables := table.GetAll()
			for name, tab := range tables {
				module := &Module{
					TablePlugin: tab,
				}
				vTab := &VirtualTable{
					TablePlugin: tab,
				}
				vTab.Cursor = &Cursor{
					TablePlugin: tab,
				}
				module.VirtualTable = vTab
				err := conn.CreateModule(name, module)
				if err != nil {
					return err
				}
				statement := table.ColumnDefinition(tab.Columns())
				format := "CREATE VIRTUAL TABLE temp." + name + " USING " + name + statement
				var values []driver.Value
				_, err = conn.Exec(format, values)
				if err != nil {
					return err
				}
			}
			return nil
		},
	})
}
