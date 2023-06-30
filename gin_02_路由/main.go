package main

import (
	"gin_02/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 配置模板文件
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "我是get方法")
	})

	// 返回json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"msg":     "你好json",
		})
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"age":  10,
			"name": "zhangsan",
		})
	})

	// 跳转前台页面并接收get传值
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "我是前台新闻数据",
		})
	})

	// 跳转前台页面并接收get传值
	r.GET("/newsGet", func(c *gin.Context) {
		// 获取前台传值
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	// 跳转不同文件夹下
	r.GET("/admin/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "我是后台首页",
			"score": 89,
			"hobby": []string{"吃饭", "睡觉", "写代码"},
		})
	})

	// Post提交数据 先跳转
	r.GET("/admin/toNews", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "sxs",
		})
	})
	// 接收post提交的数据
	r.POST("/admin/news", func(c *gin.Context) {
		username := c.PostForm("username")
		age := c.PostForm("age")
		paper := c.DefaultPostForm("paper", "111")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"paper":    paper,
		})

	})

	/*// 前台首页路由分组
	defaultRouters := r.Group("/qiantai")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "首页")
		})

		defaultRouters.GET("/admin", func(c *gin.Context) {
			c.String(200, "新闻")
		})

	}

	// 后台路由分组
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "我是后台的首页")
		})
		adminRouters.GET("/userList", func(c *gin.Context) {
			c.String(200, "后台用户列表")
		})
	}

	// api路由分组
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "我是Api的首页")
		})
		apiRouters.GET("/userList", func(c *gin.Context) {
			c.String(200, "Api用户列表")
		})
	}*/

	routers.DefaultRouters(r)
	routers.AdminRouters(r)
	routers.ApiRouters(r)

	// 启动服务
	r.Run(":8000")
}
