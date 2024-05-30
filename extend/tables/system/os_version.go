package system

import (
	"github.com/dean2021/sysql/table"
)

func GenOSVersion(context *table.QueryContext) (table.TableRows, error) {
	return table.TableRows{
		genOSVersionImpl(context),
	}, nil
}
