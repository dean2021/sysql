package sysql

import (
	"github.com/dean2021/go-sqlite3"
	"github.com/dean2021/sysql/table"
)

type VirtualTable struct {
	TablePlugin table.Table
	Cursor      sqlite3.VTabCursor
}

func (v *VirtualTable) Open() (sqlite3.VTabCursor, error) {
	return v.Cursor, nil
}
func (v *VirtualTable) BestIndex(nConstraint []sqlite3.InfoConstraint, obl []sqlite3.InfoOrderBy) (*sqlite3.IndexResult, error) {
	var used = make([]bool, len(nConstraint))
	columns := v.TablePlugin.Columns()
	for i, constraintInfo := range nConstraint {
		if !constraintInfo.Usable {
			continue
		}
		used[i] = true

		// Lookup the column name given an index into the table column set.
		if constraintInfo.Column < 0 || constraintInfo.Column >= len(columns) {
			continue
		}

		columnType := columns[constraintInfo.Column].Type
		if !table.SensibleComparison(columnType, constraintInfo.Op) {
			continue
		}

		// Check if this constraint is on an index or required column.
		//columnOptions := columns[constraintInfo.Column].Options
		//if columnOptions != table.REQUIRED && (columnOptions&(table.INDEX|table.ADDITIONAL)) != 1 {
		//	// not indexed, let sqlite filter it
		//	continue
		//}

		// Add the constraint set to the table's tracked constraints.
		columnName := columns[constraintInfo.Column].Name
		v.Cursor.(*Cursor).Constraints = append(v.Cursor.(*Cursor).Constraints, table.Constraint{
			Name: columnName,
			Op:   constraintInfo.Op,
		})
	}

	return &sqlite3.IndexResult{Used: used}, nil
}
func (v *VirtualTable) Disconnect() error { return nil }
func (v *VirtualTable) Destroy() error    { return nil }
