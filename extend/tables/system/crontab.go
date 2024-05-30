//go:build linux
// +build linux

package system

import (
	"github.com/dean2021/sysql/extend/tables/common"
	"github.com/dean2021/sysql/table"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var (
	CronSearchDirs = []string{
		"/etc/cron.d/",              // system all
		"/var/at/tabs/",             // user mac:lion
		"/var/spool/cron/",          // user linux:centos
		"/var/spool/cron/crontabs/", // user linux:debian
	}
)

type Cron struct {
	Event      string
	Command    string
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
}

func CronFileParser(filePath string) ([]Cron, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var crontabArr []Cron
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		cron := Cron{}
		if line == "" {
			continue
		}
		if strings.TrimSpace(line)[0] == '#' {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		if strings.Contains(fields[0], "@") {
			cron.Event = fields[0]
			cron.Command = fields[1]
			crontabArr = append(crontabArr, cron)
		} else {
			if len(fields) >= 6 {
				cron.Minute = fields[0]
				cron.Hour = fields[1]
				cron.DayOfMonth = fields[2]
				cron.Month = fields[3]
				cron.DayOfWeek = fields[4]
				cron.Command = strings.Join(fields[5:], " ")
				crontabArr = append(crontabArr, cron)
			}
		}
	}
	return crontabArr, nil
}

func GenCrontab(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	var cronFiles []string
	for _, cronDir := range CronSearchDirs {
		_ = filepath.WalkDir(cronDir, func(filePath string, d fs.DirEntry, err error) error {
			if d != nil && d.IsDir() {
				return nil
			}
			cronFiles = append(cronFiles, filePath)
			return nil
		})
	}
	cronFiles = append(cronFiles, common.HostEtc("crontab"))
	for _, cronFile := range cronFiles {
		crontabArr, err := CronFileParser(cronFile)
		if err != nil {
			continue
		}
		for _, cron := range crontabArr {
			results = append(results, table.TableRow{
				"event":        cron.Event,
				"command":      cron.Command,
				"minute":       cron.Minute,
				"hour":         cron.Hour,
				"day_of_month": cron.DayOfMonth,
				"month":        cron.Month,
				"day_of_week":  cron.DayOfWeek,
				"path":         cronFile,
			})
		}
	}
	return results, nil
}
