package utils

import (
	"crypto/md5"
)

func EncMd5(source string) string {
	result := md5.Sum([]byte(source))
	return string(result[:])
}
