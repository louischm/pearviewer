package utils

import (
	"os"
)

func IsDirExist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func Joins(str1, str2 string) string {
	if str1 == "" {
		return str2
	} else if str2 == "" {
		return str1
	}

	if str1[len(str1)-1:] != "/" {
		str1 += "/"
	}

	if str1[0:2] != "./" {
		str1 = "./" + str1
	}

	return str1 + str2
}
