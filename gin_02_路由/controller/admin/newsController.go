package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type NewsController struct {
}

func (con NewsController) NewsList(c *gin.Context) {
	c.String(http.StatusOK, "新闻首页")
}

func (con NewsController) NewsAdd(c *gin.Context) {
	c.String(http.StatusOK, "添加新闻-add")
}
