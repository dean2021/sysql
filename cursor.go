package sysql

import (
	"github.com/dean2021/go-sqlite3"
	"github.com/dean2021/sysql/table"
	"reflect"
)

type Cursor struct {
	index       int
	rows        table.TableRows
	TablePlugin table.Table
	Constraints table.Constraints
}

func (vc *Cursor) Column(c *sqlite3.SQLiteContext, col int) error {
	columns := vc.TablePlugin.Columns()
	value := vc.rows[vc.index][columns[col].Name]
	if value != nil {
		switch value.(type) {
		case string:
			c.ResultText(value.(string))
		case uint64:
			c.ResultInt64(int64(value.(uint64)))
		case uint32:
			c.ResultInt(int(value.(uint32)))
		case uint16:
			c.ResultInt(int(value.(uint16)))
		case uint8:
			c.ResultInt(int(value.(uint8)))
		case uint:
			c.ResultInt(int(value.(uint)))
		case int64:
			c.ResultInt64(value.(int64))
		case int32:
			c.ResultInt(int(value.(int32)))
		case int16:
			c.ResultInt(int(value.(int16)))
		case int8:
			c.ResultInt(int(value.(int8)))
		case int:
			c.ResultInt(value.(int))
		case float64:
			c.ResultDouble(value.(float64))
		case bool:
			c.ResultBool(value.(bool))
		case []byte:
			c.ResultBlob(value.([]byte))
		default:
			panic(reflect.TypeOf(value).Kind().String() + " type is not supported")
		}
	} else {
		c.ResultNull()
	}
	return nil
}

func (vc *Cursor) Filter(idxNum int, idxStr string, vals []interface{}) error {

	if len(vals) > 0 {
		for i, expr := range vals {
			vc.Constraints[i].Expr = expr
		}
	} else {
		vc.Constraints = nil
	}

	var err error
	var context = &table.QueryContext{
		Constraints: vc.Constraints,
	}
	vc.rows, err = vc.TablePlugin.Generate(context)
	if err != nil {
		return err
	}
	vc.index = 0
	return nil
}

func (vc *Cursor) Next() error {
	vc.index++
	return nil
}

func (vc *Cursor) EOF() bool {
	return vc.index >= len(vc.rows)
}

func (vc *Cursor) Rowid() (int64, error) {
	return int64(vc.index), nil
}

func (vc *Cursor) Close() error {
	return nil
}
