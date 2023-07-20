package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Identity   string `json:"identity"`    // 问题表的唯一标识
	CategoryId string `json:"category_id"` // 分类id
	Title      string `json:"title"`
	Content    string `json:"content"`
	MaxMem     int    `json:"max_mem"`     // 最大运行时间
	MaxRuntime int    `json:"max_runtime"` // 最大运行内存
}

func (table Problem) TableName() string {
	return "problem"
}
