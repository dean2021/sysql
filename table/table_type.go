package table

import "github.com/dean2021/go-sqlite3"

const (
	UNKNOWN_TYPE int = iota
	TEXT_TYPE
	INTEGER_TYPE
	BIGINT_TYPE
	UNSIGNED_BIGINT_TYPE
	DOUBLE_TYPE
	BLOB_TYPE
)

var ColumnTypeNames = map[int]string{
	UNKNOWN_TYPE:         "UNKNOWN",
	TEXT_TYPE:            "TEXT",
	INTEGER_TYPE:         "INTEGER",
	BIGINT_TYPE:          "BIGINT",
	UNSIGNED_BIGINT_TYPE: "UNSIGNED BIGINT",
	DOUBLE_TYPE:          "DOUBLE",
	BLOB_TYPE:            "BLOB",
}

func SensibleComparison(columnType int, op sqlite3.Op) bool {
	if columnType == TEXT_TYPE {
		if op == GREATER_THAN || op == GREATER_THAN_OR_EQUALS || op == LESS_THAN ||
			op == LESS_THAN_OR_EQUALS {
			return false
		}
	}
	return true
}
