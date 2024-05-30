package table

import "github.com/dean2021/go-sqlite3"

// ConstraintOperator
const (
	EQUALS                 = 2
	GREATER_THAN           = 4
	LESS_THAN_OR_EQUALS    = 8
	LESS_THAN              = 16
	GREATER_THAN_OR_EQUALS = 32
	MATCH                  = 64
	LIKE                   = 65
	GLOB                   = 66
	REGEXP                 = 67
	UNIQUE                 = 1
)

type Pair struct {
}

type ConstraintSet []Constraint

type Constraint struct {
	Name string
	Op   sqlite3.Op
	Expr interface{}
}
