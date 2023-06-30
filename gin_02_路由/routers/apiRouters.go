package routers

import "github.com/gin-gonic/gin"

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "我是Api的首页")
		})
		apiRouters.GET("/userList", func(c *gin.Context) {
			c.String(200, "Api用户列表")
		})
	}
}
