package utils

import (
	"strconv"
	"time"
)

func GetDateId() int {
	dateId, _ := strconv.Atoi(time.Now().Format("20060102"))
	return dateId
}

func GetIntPointer(n int) *int {
	return &n
}

func GetInt8Pointer(n int8) *int8 {
	return &n
}
