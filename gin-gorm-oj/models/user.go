package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
}

func (table *User) TableName() string {
	return "user"
}
