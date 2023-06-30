package middewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 逻辑操作
func InitMiddleweare(c *gin.Context) {
	// 判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println("I am InitMiddleweare中间件")

	// 保存数据
	c.Set("userName", "zhangsan")

	// 定义一个goroutine统计日志   非阻塞的，下面的程序可以正常执行
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done in path:" + c.Request.URL.Path)
	}()
}
