package utils

import (
	"github.com/louischm/logger"
	"os"
)

var log = logger.NewLog()

func IsDirExist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsFileInDir(fileName, dirName string) bool {
	if !IsDirExist(dirName) {
		return false
	}

	entries, err := os.ReadDir(dirName)
	if err != nil {
		log.Error(err.Error())
	}

	for _, entry := range entries {
		if entry.Name() == fileName {
			return true
		}
	}
	return false
}

func CreateEmptyFile(fileName, dirName string) {
	_, err := os.OpenFile(dirName+fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Error(err.Error())
	}
}
