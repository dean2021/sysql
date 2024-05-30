//go:build linux
// +build linux

package system

import (
	"github.com/dean2021/sysql/table"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	ShellHistoryFiles = []string{
		".bash_history",
		".zsh_history",
		".zhistory",
		".history",
		".sh_history",
	}
	BashTimestampRx = regexp.MustCompile("^#([0-9]+)$")
	ZshTimestampRx  = regexp.MustCompile("^: {0,10}([0-9]{1,11}):[0-9]+;(.*)$")
)

func genShellHistoryForUser(uid string, gid string, directory string) table.TableRows {
	var results table.TableRows
	for _, hfile := range ShellHistoryFiles {
		historyFile := filepath.Join(directory, hfile)
		rows, err := genShellHistoryFromFile(uid, historyFile)
		if err != nil {
			continue
		}
		for _, row := range rows {
			results = append(results, row)
		}
	}
	return results
}

func genShellHistoryFromFile(uid string, historyFile string) (table.TableRows, error) {
	var results table.TableRows
	b, err := ioutil.ReadFile(historyFile)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	total := len(lines)
	for n, line := range lines {
		var time int64
		var command string
		match := BashTimestampRx.FindStringSubmatch(line)
		if len(match) > 1 {
			if n < total {
				time, _ = strconv.ParseInt(strings.TrimSpace(match[1]), 10, 64)
				command = lines[n+1]
			}
		}
		if time == 0 {
			match := ZshTimestampRx.FindStringSubmatch(line)
			if len(match) > 1 {
				if n < total {
					time, _ = strconv.ParseInt(strings.TrimSpace(match[1]), 10, 64)
					command = lines[n+1]
				}
			} else {
				command = line
			}
		}
		if command == "" {
			continue
		}
		results = append(results, table.TableRow{
			"uid":          uid,
			"time":         time,
			"command":      command,
			"history_file": historyFile,
		})
	}
	return results, nil
}

func GenShellHistory(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	users, err := GenUsers(context)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		rows := genShellHistoryForUser(user["uid"].(string), user["gid"].(string), user["directory"].(string))
		for _, row := range rows {
			results = append(results, row)
		}
	}
	return results, nil
}
