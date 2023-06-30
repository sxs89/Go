package bootstrap

import (
	"gin_01/routes"
	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(route *gin.Engine) {

	// 注册api路由
	routes.RegisterAPIRoutes(route)

}
