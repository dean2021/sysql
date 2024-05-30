package networking

import (
	"github.com/dean2021/sysql/table"
)

func GenNetstat(context *table.QueryContext) (table.TableRows, error) {
	return genNetstatImpl(context)
}
