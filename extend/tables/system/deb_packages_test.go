package system

import (
	"github.com/dean2021/sysql/table"
	"testing"
)

func Test_parseBlock(t *testing.T) {

	var block = `
Package: zstd
Status: install ok installed
Priority: optional
Section: utils
Installed-Size: 1447
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Architecture: arm64
Source: libzstd
Version: 1.4.8+dfsg-3build1
Depends: libc6 (>= 2.34), libgcc-s1 (>= 3.3.1), liblz4-1 (>= 0.0~r127), liblzma5 (>= 5.1.1alpha+20120614), libstdc++6 (>= 12), zlib1g (>= 1:1.1.4)
Description: test
11111
Homepage: https://github.com/facebook/zstd
Original-Maintainer: Debian Med Packaging Team <debian-med-packaging@lists.alioth.debian.org>
`
	type args struct {
		block string
	}
	tests := []struct {
		name string
		args args
		want table.TableRow
	}{
		{
			name: "test1",
			args: args{
				block: block,
			},
			want: table.TableRow{
				"name":                "zstd",
				"status":              "install ok installed",
				"priority":            "optional",
				"section":             "utils",
				"size":                "1447",
				"maintainer":          "Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>",
				"arch":                "arm64",
				"source":              "libzstd",
				"version":             "1.4.8+dfsg-3build1",
				"depends":             "libc6 (>= 2.34), libgcc-s1 (>= 3.3.1), liblz4-1 (>= 0.0~r127), liblzma5 (>= 5.1.1alpha+20120614), libstdc++6 (>= 12), zlib1g (>= 1:1.1.4)",
				"summary":             " test11111",
				"homepage":            "https://github.com/facebook/zstd",
				"original_maintainer": "Debian Med Packaging Team <debian-med-packaging@lists.alioth.debian.org>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for k, v := range parseBlock(tt.args.block) {
				if val, ok := tt.want[k]; ok {

					if v.(string) != val.(string) {
						t.Errorf("parseBlock() = %v, want %v", v, val)
					}
				} else {
					t.Errorf("未知的键值: %v", k)
				}
			}
		})
	}

}
