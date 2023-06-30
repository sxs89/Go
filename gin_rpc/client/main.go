package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1.通过rpc.dial和rpc微服务端建立连接
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}

	// 2.当客户端退出的时候关闭连接
	defer conn.Close()

	// 3.调用远程函数
	var reply string
	err2 := conn.Call("hello.SayHello", "aaaa", &reply)
	if err2 != nil {
		fmt.Println(err2)
	}

	// 4.获取微服务server中的值
	fmt.Println(reply)
}
