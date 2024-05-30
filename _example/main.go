package main

import (
	"database/sql"
	"fmt"
	"github.com/dean2021/sysql"
)

func main() {

	sysql.Initialize()

	db, err := sql.Open(sysql.DriverName, ":memory:")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select pid,name,cmdline from processes")
	if err != nil {
		panic(err)
	}

	var pid string
	var name string
	var cmdline string
	for rows.Next() {
		rows.Scan(&pid, &name, &cmdline)
		fmt.Println(pid, name, cmdline)
	}
	rows.Close()
}
