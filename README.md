# sysql

>  SQL driven operating system queries can query all content on the operating system.

### Code Example
```go

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

```

## Build

go build -tags=sqlite_vtable

## Playground
```sql
-- Query all supported tables

SELECT table_name FROM information_schema.tables GROUP BY table_name;
-- Query what fields a certain table has

PRAGMA table_info('time');
-- Query what built-in functions there are

SELECT * FROM sqlite_master WHERE type='function';
-- Check if a certain process is running with root privileges, which poses a security risk

SELECT * FROM processes WHERE name LIKE '%mysql%' AND uid = 0;
SELECT * FROM processes WHERE name = 'java' AND uid = 0;
-- Find processes that delete themselves

SELECT * FROM processes WHERE on_disk = 0;
-- Determine if there are malicious commands in bash history

SELECT * FROM shell_history WHERE command LIKE '%nmap%';
-- View processes launched through a pseudo-terminal

SELECT pid, username, name, terminal FROM processes WHERE terminal != '';
-- Detect reverse shells

SELECT p.* FROM processes AS p LEFT OUTER JOIN netstat_diag AS n ON p.pid = n.pid WHERE p.name IN ('sh', 'bash', 'nc') AND n.status = 'ESTABLISHED';
-- Determine if a certain file exists
SELECT file_exists('/etc/passwd');
```
More: https://github.com/teoseller/osquery-attck

## TODO
1. psutil for sysql
2. Add more table to sysql

## Thanks

Thanks for Facebook's osquery idea


