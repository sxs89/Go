package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 定义权限模型
type Permission struct {
	ID          int
	Name        string
	Description string // 描述
}

// 定义角色模型
type Role struct {
	ID          int
	Name        string
	Description string       // 描述
	Permissions []Permission // 权限
}

// 定义用户模型
type User struct {
	ID       int
	Username string
	Password string
	Roles    []Role
}

var (
	permissions = []Permission{
		{ID: 1, Name: "create", Description: "Create resource"},
		{ID: 2, Name: "read", Description: "Read resource"},
		{ID: 3, Name: "update", Description: "Update resource"},
		{ID: 4, Name: "delete", Description: "Delete resource"},
	}

	roles = []Role{
		{
			ID: 1, Name: "admin", Description: "Administrator",
			Permissions: []Permission{permissions[0], permissions[1], permissions[2], permissions[3]},
		},
		{
			ID: 2, Name: "user", Description: "Regular user",
			Permissions: []Permission{permissions[1]},
		},
	}

	users = []User{
		{ID: 1, Username: "admin", Password: "admin123", Roles: []Role{roles[0]}},
		{ID: 2, Username: "user", Password: "user123", Roles: []Role{roles[1]}},
	}
)

// 验证用户是否具有指定权限的中间件
func authorize(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := 1

		// 查找用户
		user := findUserByID(userID)
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 检查用户是否具有指定的权限
		for _, role := range user.Roles {
			for _, p := range role.Permissions {
				if p.Name == permission {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
	}
}

// 根据ID查找用户
func findUserByID(userID int) *User {
	for _, user := range users {
		if user.ID == userID {
			return &user
		}
	}
	return nil
}

func main() {

	router := gin.Default()

	// 路由1，无需权限验证
	public := router.Group("/public")
	public.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello public"})
	})

	// 路由2， 需要权限验证
	private := router.Group("/private")
	private.Use(authorize("create"))
	private.GET("/resource", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access to private resource!"})
	})

	log.Fatal(router.Run(":8001"))

}
