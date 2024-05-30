package system

import (
	"encoding/xml"
	"github.com/dean2021/sysql/extend/tables/host"
	"github.com/dean2021/sysql/table"
	"io/ioutil"
	"strings"
)

type Plist struct {
	XMLName xml.Name `xml:"plist"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Dict    struct {
		Text   string   `xml:",chardata"`
		Key    []string `xml:"key"`
		String []string `xml:"string"`
	} `xml:"dict"`
}

func parserPListFile(path string) (map[string]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	plist := Plist{}
	err = xml.Unmarshal(data, &plist)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for k, v := range plist.Dict.Key {
		result[v] = plist.Dict.String[k]
	}
	return result, nil
}

func genOSVersionImpl(context *table.QueryContext) table.TableRow {

	var r = table.TableRow{}
	r["name"] = "macOS"
	r["major"] = "0"
	r["minor"] = "0"
	r["patch"] = "0"
	r["pid_with_namespace"] = "0"

	arch, _ := host.KernelArch()
	r["platform"] = "darwin"
	r["platform_like"] = "darwin"
	r["arch"] = arch
	kVersionPath := "/System/Library/CoreServices/SystemVersion.plist"
	plist, _ := parserPListFile(kVersionPath)
	if plist != nil {
		for key, value := range plist {
			if key == "ProductBuildVersion" {
				r["build"] = value
			} else if key == "ProductVersion" {
				r["version"] = value
			} else if key == "ProductName" {
				r["name"] = value
			}
		}
	}

	version := strings.Split(r["version"].(string), ".")
	switch len(version) {
	case 3:
		r["patch"] = version[2]
		r["minor"] = version[1]
		r["major"] = version[0]
	case 2:
		r["minor"] = version[1]
		r["major"] = version[0]
	case 1:
		r["major"] = version[0]
	}

	return r
}
