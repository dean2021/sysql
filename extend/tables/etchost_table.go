package tables

import (
	"bufio"
	"github.com/dean2021/sysql/extend/tables/common"
	"github.com/dean2021/sysql/table"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func init() {
	t := &EtcHostsTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type EtcHostsTable struct{}

func (p *EtcHostsTable) Name() string {
	return "etc_hosts"
}

func (p *EtcHostsTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "address", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "hostnames", Type: table.TEXT_TYPE, Options: table.DEFAULT},
	}
}

func (p *EtcHostsTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	var kEtcHosts string
	//	var kEtcHostsIcs string
	if runtime.GOOS != "windows" {
		kEtcHosts = common.HostEtc("hosts")
	} else {
		kEtcHosts = filepath.Join(os.Getenv("windir"), "system32\\drivers\\etc\\hosts")
		// TODO
		//	kEtcHostsIcs = filepath.Join(os.Getenv("windir"), "system32\\drivers\\etc\\hosts.ics")
	}

	f, err := os.Open(kEtcHosts)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)
	var rows table.TableRows
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 && line[0] != '#' {
			s := strings.ReplaceAll(line, "\t", " ")
			split := strings.Split(strings.TrimSpace(s), " ")
			if len(split) >= 2 {
				rows = append(rows, map[string]interface{}{
					"address":   strings.TrimSpace(split[0]),
					"hostnames": strings.TrimSpace(split[1]),
				})
			}
		}
	}

	return rows, nil
}
