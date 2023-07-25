package test

import (
	"fmt"
	"gin-gorm-oj/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.Problem, 0)
	error := db.Find(&data).Error
	if error != nil {
		t.Fatal(error)
	}

	// 打印数据
	for _, v := range data {
		fmt.Println("problem ==》 %v \n", v)
	}
}
