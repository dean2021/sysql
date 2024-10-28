//go:build windows
// +build windows

package system

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/dean2021/sysql/misc/windows"
	"github.com/dean2021/sysql/table"
	"github.com/yusufpapurcu/wmi"
)

type Win32UserAccount struct {
	AccountType        int64  `json:"accountType"`
	Caption            string `json:"caption"`
	Description        string `json:"description"`
	Disabled           bool   `json:"disabled"`
	Domain             string `json:"domain"`
	FullName           string `json:"fullName"`
	InstallDate        string `json:"installDate"`
	LocalAccount       bool   `json:"localAccount"`
	Lockout            bool   `json:"lockout"`
	PasswordChangeable bool   `json:"passwordChangeable"`
	PasswordExpires    bool   `json:"passwordExpires"`
	PasswordRequired   bool   `json:"passwordRequired"`
	Name               string `json:"name"`
	SID                string `json:"sid"`
	SIDType            int64  `json:"sidType"`
	Status             string `json:"status"`
}

type NewWin32UserAccount struct {
	Win32UserAccount `json:",inline"`
	PasswordLastSet  string `json:"passwordLastSet"`
}

func getWin32UserAccount() ([]NewWin32UserAccount, error) {
	var accounts []Win32UserAccount
	err := wmi.Query("SELECT * FROM Win32_UserAccount WHERE LocalAccount=True", &accounts)
	if err != nil {
		return nil, err
	}
	var results []NewWin32UserAccount
	for _, account := range accounts {
		results = append(results, NewWin32UserAccount{Win32UserAccount: account, PasswordLastSet: passwordLastSet(account.Name)})
	}
	return results, nil
}

func passwordLastSet(name string) string {
	c := exec.Command("cmd", "/C", fmt.Sprintf("net user %s | findstr 上次设置密码", name))
	data, err := c.CombinedOutput()
	if err != nil {
		return ""
	}
	data, err = windows.DecodeUTF16(data)
	if err != nil {
		return ""
	}
	data = bytes.TrimPrefix(data, []byte("上次设置密码"))
	return strings.TrimSpace(string(data))
}

func GenUsers(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	accounts, err := getWin32UserAccount()
	if err != nil {
		return nil, err
	}
	for _, a := range accounts {
		results = append(results, table.TableRow{
			"accountType":        a.AccountType,
			"caption":            a.Caption,
			"description":        a.Description,
			"disabled":           a.Disabled,
			"domain":             a.Domain,
			"fullName":           a.FullName,
			"installDate":        a.InstallDate,
			"localAccount":       a.LocalAccount,
			"lockout":            a.Lockout,
			"passwordChangeable": a.PasswordChangeable,
			"passwordExpires":    a.PasswordExpires,
			"passwordRequired":   a.PasswordRequired,
			"name":               a.Name,
			"sid":                a.SID,
			"sidType":            a.SIDType,
			"status":             a.Status,
			"passwordLastSet":    a.PasswordLastSet,
		})
	}
	return results, nil
}
