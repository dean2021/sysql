package system

import (
	"bufio"
	"github.com/dean2021/sysql/misc/array"
	"github.com/dean2021/sysql/table"
	"os/exec"
	"strings"
	"time"
)

func CommandWithCallback(callback func(line string), name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil {
			break
		}
		callback(line)
		time.Sleep(10 * time.Millisecond)
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

func GenRpmPackages(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows
	err := CommandWithCallback(func(line string) {
		split := strings.Split(line, "|")
		results = append(results, table.TableRow{
			"name":        array.Get(split, 0),
			"version":     array.Get(split, 1),
			"release":     array.Get(split, 2),
			"description": array.Get(split, 3),
		})
	}, "rpm", "--queryformat", "%{NAME}|%{VERSION}|%{RELEASE}|%{SUMMARY}\n", "--nosignature", "--nodigest", "--noscript", "--nofiles", "--nofiledigest", "-qa")
	return results, err
}
