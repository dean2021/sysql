package table

type TableColumns []Column

type Column struct {
	Name    string
	Type    int
	Options int
}
