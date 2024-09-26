package windows

import (
	"strconv"
	"strings"
	"time"
)

// 解决windows获取时间参数\\/Date(1718183706363)\\/的序列化问题
type Time struct {
	time.Time
}

// UnmarshalJSON for parsing non-standard JSON dates from PowerShell
func (ct *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	t, err := parseJSONDate(s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct *Time) String() string {
	if ct.IsZero() {
		return ""
	}
	return ct.Time.Format("2006-01-02 15:04:05")
}

// parseJSONDate parses non-standard JSON date format used in PowerShell outputs
func parseJSONDate(jsonDate string) (time.Time, error) {
	// Remove the escape characters and other non-numeric characters before parsing
	trim := strings.TrimPrefix(jsonDate, "\\/Date(")
	trim = strings.TrimSuffix(trim, ")\\/")
	millis, err := strconv.ParseInt(trim, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(0, millis*int64(time.Millisecond)), nil
}
