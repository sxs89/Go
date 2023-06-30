package util

import "time"

// 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取当前日期
func GetDate() string {
	template := "2023-02-03 10:59:59"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	temolate := "20061212"
	return time.Now().Format(temolate)
}
