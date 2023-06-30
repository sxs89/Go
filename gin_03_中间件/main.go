package main

import (
	"fmt"
	"gin_03/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 创建一个中间件
func initMiddleware(c *gin.Context) {
	fmt.Println("我是一个中间件")
}

func main() {

	// 创建一个默认的路由引擎
	r := gin.Default()
	// 自定义文件模板：注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixTotime": models.UnixTotime,
	})
	// 加载模板-放在路由前面
	r.LoadHTMLGlob("templates/**/*")

	// 配置静态web目录，第一个参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")

	r.Use(initMiddleware)

	//initMiddleware 表示的就是中间件
	r.GET("/", initMiddleware, func(c *gin.Context) {
		c.String(http.StatusOK, "gin的首页")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "新闻页面")
	})

	r.Run(":8001")

}
