package functions

import "testing"

func TestMd5File(t *testing.T) {

	hash := Md5File("/etc/passwd")
	if hash == "" {
		t.Fatal()
	}
	t.Logf(hash)
}
