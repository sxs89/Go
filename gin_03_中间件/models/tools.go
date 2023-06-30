package models

import (
	"fmt"
	"time"
)

// 时间戳转换成日期
func UnixTotime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2023-02-02 15:06:13")
}
