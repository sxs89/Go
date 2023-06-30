package main

import (
	"fmt"
	"gin_04/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

func main() {
	fmt.Println("文件上传")

	r := gin.Default()
	// 配置模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"message": "开始上传",
		})
	})

	r.POST("/doUpload", func(c *gin.Context) {
		// 业务处理
		userName := c.PostForm("userName")
		fmt.Println(userName)
		// 1.获取上传文件
		file, err := c.FormFile("face")
		// 2、获取后缀名，判断类型是否正确 .jpg .png .gif .jpeg
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
			".csv":  true,
		}
		_, ok := allowExtMap[extName]
		if !ok {
			c.String(200, "上传的类型不合法！")
			return
		}
		// 3.创建图片保存的目录 static/uupload/20230203
		day := util.GetDay()
		fmt.Println(day)
		dir := "./static/upload/" + day
		errs := os.MkdirAll(dir, 0666)
		if errs != nil {
			fmt.Println("创建文件失败！")
			c.String(200, "创建文件失败！")
			return
		}
		// 4、生成文件名称和文件保存目录
		unix := util.GetUnix()
		fileName := strconv.FormatInt(unix, 10) + extName

		dst := path.Join(dir, fileName)
		//dst := path.Join("./static/upload/", file.Filename)
		if err == nil {
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"username": userName,
			"dst":      dst,
		})
	})

	r.Run(":8002")
}
