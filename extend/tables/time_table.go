package tables

import (
	"github.com/dean2021/sysql/table"
	"time"
)

func init() {
	t := &TimeTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type TimeTable struct{}

func (p *TimeTable) Name() string {
	return "time"
}

func (p *TimeTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "weekday", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "year", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "month", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "day", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "hour", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "minutes", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "seconds", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "timezone", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "local_timezone", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "unix_time", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "timestamp", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "datetime", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "iso_8601", Type: table.TEXT_TYPE, Options: table.DEFAULT},

		// TODO
		//{Name: "win_timestamp", Type: table.BIGINT_TYPE, Options: table.HIDDEN},
	}
}

func (p *TimeTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	var rows table.TableRows
	now := time.Now()
	zone, _ := now.Zone()
	rows = append(rows, map[string]interface{}{
		"weekday":        now.Weekday().String(),
		"month":          now.Month().String(),
		"day":            now.Day(),
		"hour":           now.Hour(),
		"minutes":        now.Minute(),
		"seconds":        now.Second(),
		"timezone":       "UTC",
		"local_timezone": zone,
		"year":           now.Year(),
		"unix_time":      now.Unix(),
		"timestamp":      now.UTC().String(),
		"datetime":       now.Format(time.RFC3339),
		"iso_8601":       now.Format(time.RFC3339),
	})
	return rows, nil
}
