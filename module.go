package sysql

import (
	"github.com/dean2021/go-sqlite3"
	"github.com/dean2021/sysql/table"
)

type Module struct {
	VirtualTable *VirtualTable
	TablePlugin  table.Table
}

func (m *Module) Create(c *sqlite3.SQLiteConn, args []string) (sqlite3.VTab, error) {
	statement := table.ColumnDefinition(m.TablePlugin.Columns())
	format := "CREATE TABLE " + m.TablePlugin.Name() + statement
	err := c.DeclareVTab(format)
	if err != nil {
		return nil, err
	}
	return m.VirtualTable, nil
}
func (m *Module) Connect(c *sqlite3.SQLiteConn, args []string) (sqlite3.VTab, error) {
	return m.Create(c, args)
}
func (m *Module) DestroyModule() {}
