package tables

import (
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &SchemaTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type SchemaTable struct{}

func (p *SchemaTable) Name() string {
	return "schema"
}

func (p *SchemaTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "table_name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "column_name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "column_type", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "column_options", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *SchemaTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	var rows table.TableRows
	for _, tab := range table.GetAll() {
		for _, column := range tab.Columns() {
			rows = append(rows, map[string]interface{}{
				"table_name":     tab.Name(),
				"column_name":    column.Name,
				"column_type":    table.ColumnTypeNames[column.Type],
				"column_options": table.ColumnOptionsNames[column.Options],
			})
		}

	}
	return rows, nil
}
