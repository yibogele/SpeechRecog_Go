package util

import (
	CN "golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByteToString(byte []byte, charset Charset) (str string) {
	switch charset {
	case GB18030:
		var decBytes, _ = CN.GB18030.NewDecoder().Bytes(byte)
		str = string(decBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return
}

