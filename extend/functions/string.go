package functions

import (
	"fmt"
	"strings"
)

func Split(s interface{}, token interface{}, index int) string {
	split := strings.Split(fmt.Sprintf("%v", s), fmt.Sprintf("%v", token))
	if index >= len(split) {
		return ""
	}
	return split[index]
}
