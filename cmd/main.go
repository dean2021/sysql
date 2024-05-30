package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dean2021/sysql"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

var (
	App *cli.App
)

func init() {
	App = cli.NewApp()
	App.Name = "sysql cli"
	App.Usage = "sysql -h"
	App.Authors = []cli.Author{
		{Name: "Dean", Email: "dean@csoio.com"},
	}
	App.Version = "1.0.0"
	App.Copyright = "github.com/dean2021/sysql"
	App.Description = "Sysql is a powerful system query tool. Everything can be queried."
}

func isValidOutputFormat(of string) bool {
	for _, format := range []string{
		"table",
		"json",
	} {
		if of == format {
			return true
		}
	}
	return false
}

func outputTable(q string) error {
	db, err := sql.Open(sysql.DriverName, ":memory:")
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := db.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(columns)

	var values = make([]interface{}, len(columns))
	for key, _ := range values {
		var v interface{}
		values[key] = &v
	}

	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			panic(err)
		}
		var r []string
		for i, _ := range columns {
			var rawValue = *(values[i].(*interface{}))
			r = append(r, fmt.Sprintf("%v", rawValue))
		}
		table.Append(r)
	}
	defer rows.Close()

	table.Render()

	return nil
}

func outputJSON(q string) (string, error) {
	db, err := sql.Open(sysql.DriverName, ":memory:")
	if err != nil {
		return "", err
	}
	defer db.Close()
	rows, err := db.Query(q)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return "", err
		}
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func main() {

	sysql.Initialize()

	App.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "sysql, q",
			Usage:    "Execute an SQL statement",
			Required: true,
		},
		cli.StringFlag{
			Name:  "output-format, format",
			Usage: "Output format (table/json)",
			Value: "table",
		},
		cli.StringFlag{
			Name:  "output-file, file",
			Usage: "Output to file",
		},
	}

	App.Action = func(c *cli.Context) error {

		q := c.String("sysql")
		format := c.String("output-format")
		f := c.String("output-file")
		if !isValidOutputFormat(format) {
			return errors.New("ERR: output format is limited to [table / json]")
		}

		var err error
		switch format {
		case "table":
			err = outputTable(q)
			if err != nil {
				return err
			}
		case "json":
			outputStr, err := outputJSON(q)
			if err != nil {
				return err
			}
			if f == "" {
				fmt.Println(outputStr)
			} else {
				err := ioutil.WriteFile(f, []byte(outputStr), 0655)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	if err := App.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
