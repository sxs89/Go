package admin

import (
	"fmt"
	"gin_zhenghe/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserController struct {
}

type User struct {
	Id         int64      `json:"id,omitempty"`
	UserName   string     `json:"user_name,omitempty"`
	Password   string     `json:"password,omitempty"`
	CreateTime *time.Time `json:"create_time,omitempty"`
	UpdateTime *time.Time `json:"update_time,omitempty"`
}

// 创建user表
func (con UserController) InitUser(c *gin.Context) {
	errResult := utils.DB.Sync(new(User))
	if errResult != nil {
		fmt.Println("同步数据失败！")
	}
	c.String(200, "创建user表成功")
}

// 查询user表
func (con UserController) GetUSerList(c *gin.Context) {
	//userName := c.Param("username")
	//fmt.Println(userName)
	sql := "select * from user"
	userList, _ := utils.DB.QueryString(sql)
	c.JSON(200, userList)
}

/**
  根据条件查询用户信息
*/
func (con UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	var user = User{Id: 1}
	userInfo, err := utils.DB.Get(&user)
	fmt.Printf("userInfo %v", userInfo)
	fmt.Printf("error=== %v", err)
	if err != nil {
		fmt.Println("查询错误：", err)
	}
	c.JSON(200, user)
}

// 获取token
func (con UserController) GetToken(c *gin.Context) {
	userId := int64(10)
	userName := "sxs"

	tokenString, err := utils.GenerateJWT(userId, userName)
	if err != nil {
		fmt.Println("Failed to generate JWT:", err)
	}
	c.JSON(200, tokenString)

}

// 更新token过期时间
func (con UserController) UpdateJwt(c *gin.Context) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMCwidXNlcl9uYW1lIjoic3hzIiwiaXNzIjoic3hzIiwiZXhwIjoxNjg5MzE1MzAyfQ.D4Y4Zg3bICAtBWtgvgzpbSiEIjZmZe2mEr-OG_smU1M"

	newTokenString, err := utils.UpdateJwtExpiration(tokenString, time.Hour*24)
	if err != nil {
		fmt.Println("更新token过期时间失败:", err)
		return
	}
	c.JSON(200, newTokenString)
}

// 校验token
func (con UserController) ValidateJwt(c *gin.Context) {

	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMCwidXNlcl9uYW1lIjoic3hzIiwiaXNzIjoic3hzIiwiZXhwIjoxNjg5MzE1MzAyfQ.D4Y4Zg3bICAtBWtgvgzpbSiEIjZmZe2mEr-OG_smU1M"
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMCwidXNlcl9uYW1lIjoic3hzIiwiaXNzIjoic3hzIiwiZXhwIjoxNjg5MzI0MTExfQ.JlqdxgTCxSx2j-17NJwCTJ_pTFQGlSFGS5_oBy0GDm4"
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		fmt.Println("Jwt 校验失败:", err)
		return
	}
	fmt.Println(claims)
	c.JSON(200, claims)

}

// 测试着玩
func (con UserController) ApiHandler(c *gin.Context) {
	// 从上下文中获取解析出的声明信息
	claims := c.MustGet("claims").(*utils.CustomClaims)

	// 在处理函数中，可以确保 Token 校验已通过，可以进行后续操作
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("API Request Successful. Username: %s， UserID: %d", claims.UserName, claims.UserId)})
}
