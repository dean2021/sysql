package system

import (
	"bufio"
	"github.com/dean2021/sysql/table"
	"os"
	"strings"
)

var mapper = map[string]string{
	"Package":             "name",
	"Version":             "version",
	"Source":              "source",
	"Installed-Size":      "size",
	"Architecture":        "arch",
	"Status":              "status",
	"Maintainer":          "maintainer",
	"Section":             "section",
	"Priority":            "priority",
	"Depends":             "depends",
	"Description":         "description",
	"Homepage":            "homepage",
	"Original-Maintainer": "original_maintainer",
	"Replaces":            "replaces",
	"Provides":            "provides",
	"Recommends":          "recommends",
	"Suggests":            "suggests",
	"Breaks":              "breaks",
}

func isDescriptionLine(line string) bool {
	for k, _ := range mapper {
		if strings.HasPrefix(line, k+":") && k != "Description" {
			return false
		}
	}
	return true
}

func parseBlock(block string) table.TableRow {
	var row = make(table.TableRow)
	lines := strings.Split(block, "\n")
	var flag bool
	description := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "Description:") {
			flag = true
		}
		if flag {
			if isDescriptionLine(line) {
				description = description + line
			} else {
				row["description"] = strings.TrimLeft(description, "Description:")
				flag = false
			}
		}
		if !flag {
			index := strings.Index(line, ":")
			if index != -1 {
				key := line[:index]
				value := line[index+1:]
				row[mapper[key]] = strings.TrimSpace(value)
			}
		}
	}
	return row
}

func GenDebPackages(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	var pkgPath = "/var/lib/dpkg/status"
	file, err := os.Open(pkgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var block string
	var blocks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			blocks = append(blocks, block)
			block = ""
			continue
		}
		block += line + "\n"
	}
	for _, b := range blocks {
		results = append(results, parseBlock(b))
	}

	return results, nil
}
