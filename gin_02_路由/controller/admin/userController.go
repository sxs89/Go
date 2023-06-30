package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	userName, _ := c.Get("userName")
	fmt.Println(userName)
	v, _ := userName.(string)

	c.String(http.StatusOK, "用户列表"+v)
}

func (con UserController) Add(c *gin.Context) {
	c.String(http.StatusOK, "admin用户列表 -- add")
}

/*func UserAdd(c *gin.Context) {
	c.String(http.StatusOK, "用户列表-add")
}
*/
