package routes

import "github.com/gin-gonic/gin"

func HomeRoutes(r *gin.Engine) {
	homeRoutes := r.Group("/homes")
	{
		homeRoutes.GET("/", func(c *gin.Context) {
			c.String(200, "首页")
		})
	}
}
