package utils

import (
	"os"
)

func GetFileString(fileName string) (fileContent string, err error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}
