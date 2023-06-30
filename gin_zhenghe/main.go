package main

import (
	"fmt"
	"gin_zhenghe/routes"
	"gin_zhenghe/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type Student struct {
	Id        int64      `xorm:"pk autoincr" json:"id"`
	Name      string     `xorm:"unique" json:"name"`
	Age       int64      `json:"age"`
	CreatedAt *time.Time `xorm:"created" json:"created_at"`
	UpdatedAt *time.Time `xorm:"updated" json:"updated_at"`
}

func main() {
	fmt.Println("sxs")

	r := gin.Default()
	// 配置模板文件
	//r.LoadHTMLGlob("templates/**/*")
	// 加载路由
	routes.HomeRoutes(r)
	routes.AdminRoutes(r)

	// 执行
	r.Run(":8000")

}

func Init() {
	fmt.Println("main")
	err := utils.DB.Sync(new(Student))
	if err != nil {
		fmt.Println("数据库同步失败", err)
	}
}
