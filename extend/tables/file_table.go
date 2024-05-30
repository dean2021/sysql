//go:build linux || darwin

package tables

import (
	"github.com/dean2021/sysql/extend/tables/utility"
	"github.com/dean2021/sysql/table"
)

func init() {
	t := &FileTable{}
	err := table.Register(t, t.Name())
	if err != nil {
		panic(err)
	}
}

type FileTable struct{}

func (p *FileTable) Name() string {
	return "file"
}

func (p *FileTable) Columns() table.TableColumns {
	return table.TableColumns{
		{Name: "path", Type: table.TEXT_TYPE, Options: table.INDEX | table.REQUIRED},
		{Name: "directory", Type: table.TEXT_TYPE, Options: table.REQUIRED},
		{Name: "filename", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "inode", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "uid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "gid", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "mode", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		{Name: "device", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "size", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "block_size", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "atime", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "mtime", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "ctime", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "btime", Type: table.BIGINT_TYPE, Options: table.DEFAULT},
		{Name: "hard_links", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "symlink", Type: table.INTEGER_TYPE, Options: table.DEFAULT},
		{Name: "type", Type: table.TEXT_TYPE, Options: table.DEFAULT},
		//{Name: "attributes", Type: table.TEXT_TYPE, Options: table.HIDDEN},
		//{Name: "volume_serial", Type: table.TEXT_TYPE, Options: table.HIDDEN},
		//{Name: "file_id", Type: table.TEXT_TYPE, Options: table.HIDDEN},
		//{Name: "file_version", Type: table.TEXT_TYPE, Options: table.HIDDEN},
		//{Name: "product_version", Type: table.TEXT_TYPE, Options: table.HIDDEN},
		//{Name: "bsd_flags", Type: table.TEXT_TYPE, Options: table.HIDDEN},
		//{Name: "pid_with_namespace", Type: table.INTEGER_TYPE, Options: table.ADDITIONAL |table.HIDDEN},
		//{Name: "mount_namespace_id", Type: table.INTEGER_TYPE, Options: table.HIDDEN},
	}
}

func (p *FileTable) Generate(context *table.QueryContext) (table.TableRows, error) {
	rows, err := utility.GenFile(context)
	return rows, err
}
