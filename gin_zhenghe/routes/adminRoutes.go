package routes

import (
	"gin_zhenghe/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	// 后台admin路由
	adminRoutes := r.Group("/admin/")
	{
		adminRoutes.GET("/", admin.IndexController{}.Index)
		adminRoutes.GET("sxs", admin.IndexController{}.QueryUserList)

		// user路由
		adminRoutes.GET("studentInit", admin.StudentController{}.InitStudent)
		adminRoutes.GET("studentHome/*name", admin.StudentController{}.StudentHome)
		adminRoutes.GET("getByName", admin.StudentController{}.GetByName)
		adminRoutes.GET("addStudent", admin.StudentController{}.AddStudent)
		adminRoutes.GET("updateStudent", admin.StudentController{}.UpdateStudent)
		adminRoutes.GET("delteById", admin.StudentController{}.DelteById)

		// 后台user理由
		adminRoutes.GET("initUser", admin.UserController{}.InitUser)
		adminRoutes.GET("getUSerList/*username", admin.UserController{}.GetUSerList)
		adminRoutes.GET("getUserById/*id", admin.UserController{}.GetUserById)

		// Jwt token验证
		adminRoutes.GET("validateJwt", admin.UserController{}.ValidateJwt)
		adminRoutes.GET("getToken", admin.UserController{}.GetToken)
		adminRoutes.GET("updateJwt", admin.UserController{}.UpdateJwt)

		//验证jin框架
		adminRoutes.GET("/login", admin.UserController{}.ApiHandler)

		// 权限认证
		adminRoutes.GET("/initRole", admin.RoleController{}.InitRole)
		adminRoutes.GET("/queryRole", admin.RoleController{}.QueryRole)
	}

}
