package windows

import (
	"bytes"
	"io"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func DecodeUTF16(b []byte) ([]byte, error) {
	reader := transform.NewReader(ioutil.NopCloser(bytes.NewBuffer(b)), simplifiedchinese.GBK.NewDecoder())
	decoded, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return decoded, nil

}
