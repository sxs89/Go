package main

import (
	"fmt"
	"gin_zhenghe/routes"
	"gin_zhenghe/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type Student struct {
	Id        int64      `xorm:"pk autoincr" json:"id"`
	Name      string     `xorm:"unique" json:"name"`
	Age       int64      `json:"age"`
	CreatedAt *time.Time `xorm:"created" json:"created_at"`
	UpdatedAt *time.Time `xorm:"updated" json:"updated_at"`
}

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	fmt.Println("sxs")

	r := gin.Default()

	// 注册中间件，校验Token 通过这个中间件可以校验全局的token验证
	//r.Use(authMiddleware())

	// 配置模板文件
	//r.LoadHTMLGlob("templates/**/*")
	// 加载路由
	routes.HomeRoutes(r)
	routes.AdminRoutes(r)

	// 执行
	r.Run(":8000")

}

func Init() {
	fmt.Println("main")
	err := utils.DB.Sync(new(Student))
	if err != nil {
		fmt.Println("数据库同步失败", err)
	}
}

/**校验开始-------------------*/
func (con Claims) ApiHandler(c *gin.Context) {
	// 从上下文中获取解析出的声明信息
	claims := c.MustGet("claims").(*utils.CustomClaims)

	// 在处理函数中，可以确保 Token 校验已通过，可以进行后续操作
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("API Request Successful. Username: %s", claims.UserName)})
}

// 校验Token
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 检查当前请求的路径，如果在跳过 Token 验证的白名单中，则直接继续处理请求
		if isSkipAuthRoute(c.Request.URL.Path) {
			c.Next()
			return
		}

		authHeader := c.GetHeader("token")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)

		c.Next() // 放行

		//c.Abort()  // 阻断
	}
}

func isSkipAuthRoute(path string) bool {
	skipAuthRoutes := []string{
		"/public1",
		"/public2",
		"/admin/login",
	}

	for _, route := range skipAuthRoutes {
		if path == route {
			return true
		}
	}

	return false
}
