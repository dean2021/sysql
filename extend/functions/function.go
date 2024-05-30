package functions

var Functions = map[string]interface{}{
	"md5_file":    Md5File,
	"sha1_file":   Sha1File,
	"sha256_file": Sha256File,
	"md5":         Md5,
	"sha1":        Sha1,
	"sha256":      Sha256,
	"to_base64":   ToBase64,
	"from_base64": FromBase64,
	"split":       Split,
	"file_exists": FileExists,
}
