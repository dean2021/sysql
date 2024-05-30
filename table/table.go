package table

type Table interface {
	Name() string
	Columns() TableColumns
	Generate(context *QueryContext) (TableRows, error)
}

type UsedColumns []string
