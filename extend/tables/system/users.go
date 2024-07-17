//go:build linux || darwin

package system

import (
	"bufio"
	"os"
	"strings"

	"github.com/dean2021/sysql/extend/tables/common"
	"github.com/dean2021/sysql/table"
)

func GenUsers(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	fi, err := os.Open(common.HostEtc("passwd"))
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		s, _, err := br.ReadLine()
		if err != nil {
			break
		}
		line := strings.TrimSpace(string(s))
		if line != "" && line[:1] != "#" {
			items := strings.Split(line, ":")
			if len(items) < 7 {
				continue
			}
			results = append(results, table.TableRow{
				"uid":         items[2],
				"gid":         items[3],
				"username":    items[0],
				"description": items[4],
				"directory":   items[5],
				"shell":       items[6],
			})
		}
	}
	return results, nil
}
