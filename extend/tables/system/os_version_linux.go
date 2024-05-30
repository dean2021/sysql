package system

import (
	"github.com/dean2021/sysql/extend/tables/host"
	"github.com/dean2021/sysql/table"
)

func genOSVersionImpl(context *table.QueryContext) table.TableRow {
	var r = table.TableRow{}
	r["name"] = "Unknown"
	r["major"] = "0"
	r["minor"] = "0"
	r["patch"] = "0"
	r["platform"] = "posix"
	r["pid_with_namespace"] = "0"
	info, _ := host.Info()
	if info != nil {
		r["arch"] = info.KernelArch
		r["platform"] = info.Platform
		r["platform_like"] = info.Platform
		r["version"] = info.PlatformVersion
		r["name"] = info.OS
		// No build name.
		r["build"] = ""
	}
	return r
}
