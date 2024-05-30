//go:build windows

package system

import (
	"errors"
	"github.com/dean2021/sysql/table"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// CheckConnectWindowsUpdateServer 检查是否可以访问windows更新服务器
func CheckConnectWindowsUpdateServer() bool {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	url := "http://ds.download.windowsupdate.com/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	if !strings.Contains(string(bytes), "http://windowsupdate.microsoft.com") {
		return false
	}
	return true
}

// WUAPackage describes a Windows Update Agent package.
type WUAPackage struct {
	Title                    string    `json:"title"`
	Description              string    `json:"description"`
	Categories               []string  `json:"categories"`
	CategoryIDs              []string  `json:"categoryIDs"`
	KBArticleIDs             []string  `json:"kbArticleIDs"`
	MoreInfoURLs             []string  `json:"moreInfoURLs"`
	SupportURL               string    `json:"supportURL"`
	UpdateID                 string    `json:"updateID"`
	RevisionNumber           int32     `json:"revisionNumber"`
	MSRCSeverity             string    `json:"msrcSeverity"`
	LastDeploymentChangeTime time.Time `json:"lastDeploymentChangeTime"`
}

func GenWindowsUpdates(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows

	// 参考:https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher-search
	packages, err := WUAUpdates("IsInstalled=0")
	if err != nil {
		return results, errors.New("failed to obtain Windows system patch information:" + err.Error())
	}
	if packages == nil {
		return results, errors.New("error: No patch data obtained")
	}
	for _, pack := range packages {
		results = append(results, table.TableRow{
			"title":                       pack.Title,
			"description":                 pack.Description,
			"categories":                  strings.Join(pack.Categories, ","),
			"kb_article_ids":              strings.Join(pack.KBArticleIDs, ","),
			"more_info_urls":              strings.Join(pack.MoreInfoURLs, ","),
			"support_url":                 pack.SupportURL,
			"update_id":                   pack.UpdateID,
			"revision_number":             pack.RevisionNumber,
			"severity":                    pack.MSRCSeverity,
			"last_deployment_change_time": pack.LastDeploymentChangeTime.Format("2006-01-02 15:04:05"),
		})
	}
	return results, nil
}
