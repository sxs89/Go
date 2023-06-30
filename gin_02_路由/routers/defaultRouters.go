package routers

import "github.com/gin-gonic/gin"

func DefaultRouters(r *gin.Engine) {
	defaultRouters := r.Group("/qiantai")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "首页")
		})

		defaultRouters.GET("/admin", func(c *gin.Context) {
			c.String(200, "新闻")
		})

	}
}
