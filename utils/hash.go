package utils

import (
	"bytes"
	"crypto/sha256"
	"io"
	"strconv"
)

func CalculateHash(any string) string {
	hs := sha256.New()
	io.WriteString(hs, any)
	return ByteToHex(hs.Sum(nil))
}

func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}

	return buffer.String()
}
