package admin

import (
	"fmt"
	"gin_zhenghe/utils"
	"github.com/gin-gonic/gin"
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
	userName := c.Param("username")
	fmt.Println(userName)
	sql := "select * from user"
	userList, _ := utils.DB.QueryString(sql)
	c.JSON(200, userList)
}
