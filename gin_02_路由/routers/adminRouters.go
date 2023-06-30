package routers

import (
	"gin_02/controller/admin"
	"gin_02/middewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/admin", middewares.InitMiddleweare)
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "我是后台的首页")
		})

		adminRouters.GET("/message", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin/message.html", gin.H{
				"message": "后台信息页",
			})
		})

		adminRouters.GET("/userList", func(c *gin.Context) {
			c.String(200, "后台用户列表")
		})
		// 改写上面方案
		//adminRouters.GET("/userAdd", admin.UserAdd)

		// 继续改造
		adminRouters.GET("/indexs", admin.UserController{}.Index)
		adminRouters.GET("/userAdd", admin.UserController{}.Add)

		// 新闻路由
		adminRouters.GET("/newsList", admin.NewsController{}.NewsList)
		adminRouters.GET("/newsAdd", admin.NewsController{}.NewsAdd)
	}
}
