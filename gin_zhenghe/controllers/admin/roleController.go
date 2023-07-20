package admin

import (
	"fmt"
	"gin_zhenghe/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Users struct {
	Id       int64 `json:"id,omitempty"`
	Username string
	Password string
	Roles    []*Role `xorm:"-"` // 忽略映射关系
}

type Role struct {
	Id          int64
	Name        string
	Permissions []*Permission `xorm:"-"` // 忽略映射关系
}

type Permission struct {
	Id   int64 `json:"id,omitempty"`
	Name string
}

type RoleController struct {
}

// 创建表
func (con RoleController) InitRole(c *gin.Context) {
	errResult := utils.DB.Sync(new(Users), new(Role), new(Permission))
	if errResult != nil {
		fmt.Println("初始化数据失败！")
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// 获取role数据
func (con RoleController) QueryRole(c *gin.Context) {
	roles, err := utils.DB.QueryString("select * from g_role")
	if err != nil {
		fmt.Println("查询role数据失败", err)
	}

	for _, item := range roles {
		//fmt.Println(key, item)
		for k, ro := range item {
			fmt.Println(k, ro)
		}
	}

	//fmt.Println(roles)
	c.JSON(http.StatusOK, gin.H{
		"rols":  roles,
		"hello": "world",
	})

}
