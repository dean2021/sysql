package system

import (
	"github.com/dean2021/sysql/table"
	"github.com/shirou/gopsutil/v3/mem"
)

func GenMemoryInfo(context *table.QueryContext) (table.TableRows, error) {

	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	rows := table.TableRows{
		{
			"total":            memory.Total,
			"available":        memory.Available,
			"used":             memory.Used,
			"used_percent":     memory.UsedPercent,
			"free":             memory.Free,
			"active":           memory.Active,
			"inactive":         memory.Inactive,
			"wired":            memory.Wired,
			"laundry":          memory.Laundry,
			"buffers":          memory.Buffers,
			"cached":           memory.Cached,
			"write_back":       memory.WriteBack,
			"dirty":            memory.Dirty,
			"write_back_tmp":   memory.WriteBackTmp,
			"shared":           memory.Shared,
			"slab":             memory.Slab,
			"sreclaimable":     memory.Sreclaimable,
			"sunreclaim":       memory.Sunreclaim,
			"page_tables":      memory.PageTables,
			"swap_cached":      memory.SwapCached,
			"commit_limit":     memory.CommitLimit,
			"committed_a_s":    memory.CommittedAS,
			"high_total":       memory.HighTotal,
			"high_free":        memory.HighFree,
			"low_total":        memory.LowTotal,
			"low_free":         memory.LowFree,
			"swap_total":       memory.SwapTotal,
			"swap_free":        memory.SwapFree,
			"mapped":           memory.Mapped,
			"vmalloc_total":    memory.VmallocTotal,
			"vmalloc_used":     memory.VmallocUsed,
			"vmalloc_chunk":    memory.VmallocChunk,
			"huge_pages_total": memory.HugePagesTotal,
			"huge_pages_free":  memory.HugePagesFree,
			"huge_pages_rsvd":  memory.HugePagesRsvd,
			"huge_pages_surp":  memory.HugePagesSurp,
			"huge_page_size":   memory.HugePageSize,
			"anon_huge_pages":  memory.AnonHugePages},
	}
	return rows, nil
}
