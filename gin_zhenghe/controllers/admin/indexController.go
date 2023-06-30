package admin

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	c.String(200, "后台首页")
}

func (con IndexController) QueryUserList(c *gin.Context) {
	c.String(200, "success 史心胜")
}
