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

select table_name from schema group by table_name
-- Query what fields a certain table has

PRAGMA table_info('time');
-- Query what built-in functions there are

SELECT * FROM functions;
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
-- Ping any host
SELECT * FROM ping WHERE addr='www.google.com'
-- Returns the Listening port List - ATT&CK T1043,T1090,T1094,T1205,T1219,T1105,T1065,T1102
select p.name, p.path, lp.local_port, lp.local_address, lp.protocol  from netstat lp LEFT JOIN processes p ON lp.pid = p.pid WHERE lp.local_port != 0 AND p.name != '';
```
More: https://github.com/teoseller/osquery-attck

## Tables

### Windows tables


|   TABLE NAME    |
|-----------------|
| cpu_info        |
| curl            |
| etc_hosts       |
| functions       |
| hash            |
| interfaces      |
| last            |
| list            |
| memory_info     |
| netstat         |
| os_version      |
| patches         |
| ping            |
| processes       |
| schema          |
| services        |
| time            |
| users           |
| windows_product |
| windows_update  |

### Linux tables


|  TABLE NAME   |
|---------------|
| cpu_info      |
| crontab       |
| curl          |
| deb_packages  |
| etc_hosts     |
| file          |
| functions     |
| hash          |
| interfaces    |
| last          |
| list          |
| load_average  |
| memory_info   |
| netstat       |
| netstat_diag  |
| os_version    |
| ping          |
| processes     |
| rpm_packages  |
| schema        |
| shell_history |
| time          |
| users         |



## TODO
1. Add NPM table
2. Add Pip table
3. Add Jar table
4. Add more function to sysql


## Thanks

Thanks for Facebook's osquery idea


