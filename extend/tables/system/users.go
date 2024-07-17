//go:build linux || darwin

package system

import (
	"bufio"
	"os"
	"strings"

	"github.com/dean2021/sysql/extend/tables/common"
	"github.com/dean2021/sysql/table"
)

type User struct {
	Uid         string `json:"uid"`
	Gid         string `json:"gid"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Directory   string `json:"directory"`
	Shell       string `json:"shell"`
}

func GenUsers(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	users, err := getUsers()
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		results = append(results, table.TableRow{
			"uid":         u.Uid,
			"gid":         u.Gid,
			"username":    u.Username,
			"description": u.Description,
			"directory":   u.Directory,
			"shell":       u.Shell,
		})
	}
	return results, nil
}

func getUsers() ([]User, error) {
	fi, err := os.Open(common.HostEtc("passwd"))
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	users := []User{}
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
			users = append(users, User{
				Uid:         items[2],
				Gid:         items[3],
				Username:    items[0],
				Description: items[4],
				Directory:   items[5],
				Shell:       items[6],
			})
		}
	}
	return users, nil
}
