package table

import (
	"testing"
)

func TestColumnDefinition(t *testing.T) {

	tests := TableColumns{
		{Name: "pid", Type: BIGINT_TYPE, Options: INDEX},
		{Name: "name", Type: TEXT_TYPE, Options: DEFAULT},
		{Name: "path", Type: TEXT_TYPE, Options: DEFAULT},
		{Name: "cmdline", Type: TEXT_TYPE, Options: DEFAULT},
		{Name: "state", Type: TEXT_TYPE, Options: DEFAULT},
		{Name: "cwd", Type: TEXT_TYPE, Options: DEFAULT},
		{Name: "root", Type: TEXT_TYPE, Options: DEFAULT},
		{Name: "uid", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "gid", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "euid", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "egid", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "suid", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "sgid", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "on_disk", Type: INTEGER_TYPE, Options: DEFAULT},
		{Name: "wired_size", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "resident_size", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "total_size", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "user_time", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "system_time", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "disk_bytes_read", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "disk_bytes_written", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "start_time", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "parent", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "pgroup", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "threads", Type: INTEGER_TYPE, Options: DEFAULT},
		{Name: "nice", Type: INTEGER_TYPE, Options: DEFAULT},
		{Name: "elevated_token", Type: INTEGER_TYPE, Options: HIDDEN},
		{Name: "secure_process", Type: INTEGER_TYPE, Options: HIDDEN},
		{Name: "protection_type", Type: TEXT_TYPE, Options: HIDDEN},
		{Name: "virtual_process", Type: INTEGER_TYPE, Options: HIDDEN},
		{Name: "elapsed_time", Type: BIGINT_TYPE, Options: HIDDEN},
		{Name: "handle_count", Type: BIGINT_TYPE, Options: HIDDEN},
		{Name: "percent_processor_time", Type: BIGINT_TYPE, Options: HIDDEN},
		{Name: "upid", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "uppid", Type: BIGINT_TYPE, Options: DEFAULT},
		{Name: "cpu_type", Type: INTEGER_TYPE, Options: DEFAULT},
		{Name: "cpu_subtype", Type: INTEGER_TYPE, Options: DEFAULT},
	}

	t.Run("ColumnDefinition", func(t *testing.T) {
		definition := ColumnDefinition(tests)
		t.Log(definition)
	})

}
