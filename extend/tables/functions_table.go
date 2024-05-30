package tables

import (
	"github.com/dean2021/sysql/extend/functions"
	"github.com/dean2021/sysql/table"
	"reflect"
)

func init() {
	t := &FunctionsTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type FunctionsTable struct{}

func (p *FunctionsTable) Name() string {
	return "functions"
}

func (p *FunctionsTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "func_name", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *FunctionsTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	var rows table.TableRows
	for k, v := range functions.Functions {
		funName := k + "("
		fn := reflect.ValueOf(v)
		length := fn.Type().NumIn()
		for i := 0; i < length; i++ {
			funType := fn.Type().In(i).Name()
			if funType == "" {
				funType = "any"
			}
			if i < (length - 1) {
				funName += funType + ","
			} else {
				funName += funType
			}
		}
		funName += ")"
		rows = append(rows, map[string]interface{}{
			"func_name": funName,
		})
	}
	return rows, nil
}
