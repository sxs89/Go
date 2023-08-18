package service

import (
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProblemList(c *gin.Context) {
	models.GetProblemList()
	c.String(http.StatusOK, "get ProblemList success")
	//page := c.DefaultQuery("page", define.DefaultPage)
	//size := c.DefaultQuery("size", define.DefaultSize)
	//keyword := c.Query("keyword")
	//models.GetPro
}
