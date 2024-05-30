//go:build windows

package system

import (
	"errors"
	"github.com/dean2021/sysql/table"
	"github.com/shirou/gopsutil/v3/winservices"
	"golang.org/x/sys/windows/svc"
	"log"
)

/*
	Stopped         = State(windows.SERVICE_STOPPED)
	StartPending    = State(windows.SERVICE_START_PENDING)
	StopPending     = State(windows.SERVICE_STOP_PENDING)
	Running         = State(windows.SERVICE_RUNNING)
	ContinuePending = State(windows.SERVICE_CONTINUE_PENDING)
	PausePending    = State(windows.SERVICE_PAUSE_PENDING)
	Paused          = State(windows.SERVICE_PAUSED)
*/

func getStartType(startType uint32) string {
	kSvcStartType := []string{
		"BOOT_START", "SYSTEM_START", "AUTO_START", "DEMAND_START", "DISABLED",
	}

	log.Println(startType)
	if startType < 0 || startType > 4 {
		return "UNKNOWN"
	}

	return kSvcStartType[startType]
}

func getState(state svc.State) string {
	states := []string{
		"UNKNOWN",
		"STOPPED",
		"START_PENDING",
		"STOP_PENDING",
		"RUNNING",
		"CONTINUE_PENDING",
		"PAUSE_PENDING",
		"PAUSED",
	}
	if state < 0 || state > 7 {
		return "UNKNOWN"
	}
	return states[state]
}

func getServiceType(pid uint32) string {
	var kServiceType = map[uint32]string{
		0x00000001: "KERNEL_DRIVER",
		0x00000002: "FILE_SYSTEM_DRIVER",
		0x00000010: "OWN_PROCESS",
		0x00000020: "SHARE_PROCESS",
		0x00000050: "USER_OWN_PROCESS",
		0x00000060: "USER_SHARE_PROCESS",
		0x000000d0: "USER_OWN_PROCESS(Instance)",
		0x000000e0: "USER_SHARE_PROCESS(Instance)",
		0x00000100: "INTERACTIVE_PROCESS",
		0x00000110: "OWN_PROCESS(Interactive)",
		0x00000120: "SHARE_PROCESS(Interactive)",
	}
	return kServiceType[pid]
}

func GenWindowsServices(context *table.QueryContext) (table.TableRows, error) {
	var results table.TableRows

	// 参考:https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher-search
	services, err := winservices.ListServices()
	if err != nil {
		return results, errors.New("failed to obtain Windows services information:" + err.Error())
	}

	for _, s := range services {
		service, err := winservices.NewService(s.Name)
		if err != nil {
			continue
		}
		config, err := service.QueryServiceConfig()
		if err != nil {
			continue
		}
		status, err := service.QueryStatus()
		if err != nil {
			continue
		}
		results = append(results, table.TableRow{
			"name":            s.Name,
			"service_type":    getServiceType(config.ServiceType),
			"status":          getState(status.State),
			"pid":             status.Pid,
			"path":            config.BinaryPathName,
			"win32_exit_code": status.Win32ExitCode,
			"display_name":    config.DisplayName,
			"start_type":      getStartType(config.StartType),
		})
	}
	return results, nil
}
