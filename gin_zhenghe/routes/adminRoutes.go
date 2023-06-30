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
	}

}
