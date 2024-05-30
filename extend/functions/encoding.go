package functions

import (
	"encoding/base64"
	"fmt"
)

func ToBase64(s interface{}) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", s)))
}

func FromBase64(s interface{}) string {
	decodeString, err := base64.StdEncoding.DecodeString(fmt.Sprintf("%v", s))
	if err != nil {
		return ""
	}
	return string(decodeString)
}
