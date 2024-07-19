//go:build linux

package system

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/dean2021/sysql/extend/tables/common"
	"github.com/dean2021/sysql/table"
)

type Shadow struct {
	Username       string `json:"username"`
	PasswordStatus string `json:"password_status"`
	Hash           string `json:"hash"`        //加密后的用户密码
	LastChange     int64  `json:"last_change"` //上次密码修改时间。从1970年1月1日以来的天数
	Min            int64  `json:"min"`         //密码最短使用时间。在这段时间内用户不能更改密码。
	Max            int64  `json:"max"`         //密码最长使用时间。到达这个天数后必须更改密码。
	Warning        int64  `json:"warning"`     //密码到期前多少天发出警告
	Inactive       int64  `json:"inactive"`    //密码过期后账户被禁用前的天数。在这段时间内如果用户未更改密码，账户将被禁用。
	Expire         int64  `json:"expire"`      //账户过期日期。从1970年1月1日以来的天数，过了这个日期，账户将被完全禁用。
	Flag           int64  `json:"flag"`        //保留字段，目前未被广泛使用，通常为空。
}

func GenShadow(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	shadows, err := getShadows()
	if err != nil {
		return nil, err
	}
	for _, s := range shadows {
		results = append(results, table.TableRow{
			"username":        s.Username,
			"password_status": s.PasswordStatus,
			"hash":            s.Hash,
			"last_change":     s.LastChange,
			"min":             s.Min,
			"max":             s.Max,
			"warning":         s.Warning,
			"inactive":        s.Inactive,
			"expire":          s.Expire,
		})
	}
	return results, nil
}

func getShadows() ([]Shadow, error) {
	fi, err := os.Open(common.HostEtc("shadow"))
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	shadows := []Shadow{}
	for {
		s, _, err := br.ReadLine()
		if err != nil {
			break
		}
		line := strings.TrimSpace(string(s))
		if line != "" && line[:1] != "#" {
			items := strings.Split(line, ":")
			if len(items) < 9 {
				continue
			}
			s := Shadow{
				Username:   items[0],
				LastChange: mustParseInt64(items[2]),
				Min:        mustParseInt64(items[3]),
				Max:        mustParseInt64(items[4]),
				Warning:    mustParseInt64(items[5]),
				Inactive:   mustParseInt64(items[6]),
				Expire:     mustParseInt64(items[7]),
				Flag:       mustParseInt64(items[8]),
			}
			passwdHash := items[1]
			if passwdHash == "!!" {
				s.PasswordStatus = "not_set"
			} else if passwdHash[0] == '!' || passwdHash[0] == '*' || passwdHash[0] == 'x' {
				s.PasswordStatus = "locked"
			} else if passwdHash == "" {
				s.PasswordStatus = "empty"
			} else {
				s.PasswordStatus = "active"
			}
			s.Hash = passwdHash
			shadows = append(shadows, s)
		}
	}
	return shadows, nil
}

func mustParseInt64(str string) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return val
}
