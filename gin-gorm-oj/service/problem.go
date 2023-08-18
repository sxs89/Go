package service

import (
	"gin-gorm-oj/define"
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetProblemList(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	keyword := c.Query("keyword")
	page = (page - 1) * size
	data := make([]*models.Problem, 0)
	tx := models.GetProblemList(keyword)
	err1 := tx.Offset(page).Limit(size).Find(&data).Error
	if err1 != nil {
		log.Println("Get problem list Error", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
