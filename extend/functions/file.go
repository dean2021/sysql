package functions

import (
	"fmt"
	"github.com/dean2021/sysql/misc/os"
)

// FileExists
// @example : select name, file_exists(path) as fe from processes where fe = 0
func FileExists(path interface{}) int {
	exists, err := os.PathExists(fmt.Sprintf("%v", path))
	if err != nil {
		return -1
	}
	if exists {
		return 1
	} else {
		return 0
	}
}
