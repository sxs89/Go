package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
}

func (table *Category) TableName() string {
	return "category"
}
