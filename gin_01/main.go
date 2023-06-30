package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	fmt.Println("sxs")
	// 创建路由
	//r := gin.Default()
	// 根据路由规则，执行函数
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello world")
	//})
	//bootstrap.SetupRoute(r)
	// 监听端口
	//r.Run(":8000")

	// 异步
	r := gin.Default()
	r.GET("/long_async", func(c *gin.Context) {
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})
	r.Run(":8000")
}
