package tables

import (
	"github.com/dean2021/sysql/extend/tables/system"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &MemoryInfoTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type MemoryInfoTable struct{}

func (p *MemoryInfoTable) Name() string {
	return "memory_info"
}

func (p *MemoryInfoTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "total", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "available", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "used", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "used_percent", Type: table.DOUBLE_TYPE, Options: table.DEFAULT},
		{Name: "free", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "active", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "inactive", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "wired", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "laundry", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "buffers", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "cached", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "write_back", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "dirty", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "write_back_tmp", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "shared", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "slab", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "sreclaimable", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "sunreclaim", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "page_tables", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "swap_cached", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "commit_limit", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "committed_a_s", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "high_total", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "high_free", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "low_total", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "low_free", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "swap_total", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "swap_free", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "mapped", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "vmalloc_total", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "vmalloc_used", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "vmalloc_chunk", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "huge_pages_total", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "huge_pages_free", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "huge_pages_rsvd", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "huge_pages_surp", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "huge_page_size", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "anon_huge_pages", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
	}
}

func (p *MemoryInfoTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	list, err := system.GenMemoryInfo(context)
	if err != nil {
		return nil, err
	}
	return list, nil
}
