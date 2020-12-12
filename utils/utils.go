package utils

import "strings"

func HandleStrings(str string) string  {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)

	return str
}