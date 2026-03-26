package utils

import (
	"gin_rbac_api/global"
	"strings"
)

func SegToSearchWords(text string) (ws string) {
	wordChan := global.SYS_JIEBA.CutForSearch(text, true)
	var strs []string
	for s := range wordChan {
		strs = append(strs, s)
	}
	return strings.Join(strs, "|")
}
