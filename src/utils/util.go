package utils

import (
	"crypto/md5"
)

var ISDEBUG bool = true

func EncMd5(source string) string {
	result := md5.Sum([]byte(source))
	return string(result[:])
}
