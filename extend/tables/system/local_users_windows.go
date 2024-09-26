//go:build windows
// +build windows

package system

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/dean2021/sysql/misc/windows"
	"github.com/dean2021/sysql/table"
)

type LocalUserAccount struct {
	Name                   string       `json:"Name"`
	Description            string       `json:"Description"`
	PrincipalSource        int          `json:"PrincipalSource"`
	AccountExpires         windows.Time `json:"AccountExpires"`
	PasswordLastSet        windows.Time `json:"PasswordLastSet"`
	PasswordChangeableDate windows.Time `json:"PasswordChangeableDate"`
	PasswordExpires        windows.Time `json:"PasswordExpires"`
	UserMayChangePassword  bool         `json:"UserMayChangePassword"`
	Enabled                bool         `json:"Enabled"`
	LastLogon              windows.Time `json:"LastLogon"`
}

func getLocalUserAccount() ([]LocalUserAccount, error) {
	var s []LocalUserAccount
	cmd := exec.Command("powershell", "-Command", "Get-LocalUser | Select-Object * | ConvertTo-Json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	b, err := windows.DecodeUTF16(out)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &s)
	return s, err
}

func GenLocalUsers(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	accounts, err := getLocalUserAccount()
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	for _, a := range accounts {
		results = append(results, table.TableRow{
			"name":                   a.Name,
			"description":            a.Description,
			"principalSource":        a.PrincipalSource,
			"accountExpires":         a.AccountExpires.String(),
			"passwordLastSet":        a.PasswordLastSet.String(),
			"passwordChangeableDate": a.PasswordChangeableDate.String(),
			"passwordExpires":        a.PasswordExpires.String(),
			"userMayChangePassword":  a.UserMayChangePassword,
			"enabled":                a.Enabled,
			"lastLogon":              a.LastLogon.String(),
		})
	}
	return results, nil
}
