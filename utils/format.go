package utils

import (
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/lib/pq"
)

func GetGenderStr(gender int8) string {
	if gender == 1 {
		return "男"
	} else {
		return "女"
	}
}
func GetUserGenderStr(gender *int) string {
	if gender != nil {
		if *gender == 1 {
			return "男"
		} else {
			return "女"
		}
	} else {
		return "未知"
	}
}

func GetYearMonth(date *pgtype.Date) string {
	if date != nil {
		return date.Time.Format("2006-01")
	} else {
		return "未知"
	}
}

func GetEduLevelName(eduLevel int) string {
	switch eduLevel {
	case 1:
		return "初中及以下"
	case 2:
		return "高中"
	case 3:
		return "中专/中技"
	case 4:
		return "大专"
	case 5:
		return "本科"
	case 6:
		return "硕士"
	case 7:
		return "博士"
	}
	return "未知"
}

func GetCurrentName(current int) string {
	switch current {
	case 1:
		return "离职-随时到岗"
	case 2:
		return "在职-月内到岗"
	case 3:
		return "在职-考虑机会"
	case 4:
		return "在职-暂不考虑"
	}
	return "未知"
}

func GetScaleName(scale int) string {
	switch scale {
	case 1:
		return "50人以下"
	case 2:
		return "50-99人"
	case 3:
		return "100-499人"
	case 4:
		return "500-999人"
	case 5:
		return "1000-9999人"
	case 6:
		return "10000人以上"
	}
	return "未知"
}

func GetExpYearName(expYear int) string {
	switch expYear {
	case 1:
		return "无经验要求"
	case 2:
		return "1-3年经验"
	case 3:
		return "3-5年经验"
	case 4:
		return "5-10年经验"
	case 5:
		return "10年以上经验"
	}
	return "未知"
}

func GetPosTypeName(posType int) string {
	if posType == 1 {
		return "全职"
	} else {
		return "兼职"
	}
}

func GetRestTypeName(restType int) string {
	switch restType {
	case 1:
		return "双休"
	case 2:
		return "单休"
	case 3:
		return "大小周"
	case 4:
		return "其他"
	}
	return "未知"
}
func GetBizTypeName(bizType int) string {
	switch bizType {
	case 1:
		return "民营"
	case 2:
		return "国企"
	case 3:
		return "合资"
	case 4:
		return "外资"
	case 5:
		return "其他"
	}
	return "未知"
}

func GetStatusStr(status int8) string {
	if status == 1 {
		return "可用"
	} else {
		return "停用"
	}
}

func ConvertArrayToString(stringArray pq.StringArray, code string) string {
	tString := ""
	if stringArray != nil {
		for i := 0; i < len(stringArray); i++ {
			fmt.Println(stringArray[i])
			tString += stringArray[i] + code
		}
	}
	return tString
}
