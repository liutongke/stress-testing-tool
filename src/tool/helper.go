package tool

// Package helper 帮助函数，时间、数组的通用处理
import (
	"time"
)

// DiffNano 时间差，纳秒
func DiffNano(startTime time.Time) (diff time.Duration) {
	diff = time.Since(startTime)
	return
}

// InArrayStr 判断字符串是否在数组内
func InArrayStr(str string, arr []string) (inArray bool) {
	for _, s := range arr {
		if s == str {
			inArray = true
			break
		}
	}
	return
}
