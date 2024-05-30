package system

import "C"
import (
	"github.com/dean2021/sysql/extend/tables/host"
	"github.com/dean2021/sysql/table"
)

const (
	EMPTY         = 0
	RUN_LVL       = 1
	BOOT_TIME     = 2
	OLD_TIME      = 3
	NEW_TIME      = 4
	INIT_PROCESS  = 5
	LOGIN_PROCESS = 6
	USER_PROCESS  = 7
	DEAD_PROCESS  = 8
)

func typeNameForType(t int) string {
	switch t {
	case EMPTY:
		return "empty"
	case RUN_LVL:
		return "run-level"
	case BOOT_TIME:
		return "boot-time"
	case NEW_TIME:
		return "new-time"
	case OLD_TIME:
		return "old-time"
	case INIT_PROCESS:
		return "init-process"
	case LOGIN_PROCESS:
		return "login-process"
	case USER_PROCESS:
		return "user-process"
	case DEAD_PROCESS:
		return "dead-process"
	}
	return ""
}

func GenLast(context *table.QueryContext) (table.TableRows, error) {

	var results table.TableRows
	users, err := host.Users()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		results = append(results, table.TableRow{
			"username":  user.User,
			"tty":       user.Terminal,
			"host":      user.Host,
			"type":      user.Type,
			"type_name": typeNameForType(user.Type),
			"pid":       user.Pid,
			"time":      user.Started,
		})
	}

	return results, nil
}
