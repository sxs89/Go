package router

import (
	"gin-gorm-oj/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	// 路由规则
	r := gin.Default()
	appRouters := r.Group("/app")
	{
		appRouters.GET("/ping", service.Ping)
	}

	return r

}
