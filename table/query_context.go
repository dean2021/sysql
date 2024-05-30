package table

type QueryContext struct {
	Constraints Constraints
}

type Constraints []Constraint

func (c *Constraints) Size() int {
	return len(*c)
}
func (c *Constraints) Exists(columnName string, op int) bool {
	for _, v := range *c {
		if v.Name == columnName && int(v.Op) == op {
			return true
		}
	}
	return false
}

func (c *Constraints) GetAll(columnName string, op int) []interface{} {
	var results []interface{}
	for _, v := range *c {
		if v.Name == columnName && int(v.Op) == op {
			if v.Expr != nil {
				results = append(results, v.Expr)
			}
		}
	}
	return results
}

func (c *Constraints) Count(columnName string) int {
	num := 0
	for _, v := range *c {
		if v.Name == columnName {
			num++
		}
	}
	return num
}
