package main

import (
	"fmt"
	"net/rpc"
)

// 传输的结构体
type AddgoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

// 返回的结构体
type AddgoodsRes struct {
	Success bool
	Message string
}

// ---------------------------------------------

// 根据ID获取对应参数的结构体
type GetGoodsReq struct {
	Id int
}

// 返回过去的值
type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func main() {
	// 通过rpc建立连接
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8020")
	if err1 != nil {
		fmt.Println(err1)
	}
	// 2、客户端退出时建立连接
	defer conn.Close()

	// 3、远程调用函数
	addRequest := AddgoodsReq{
		Id:      1,
		Title:   "商品标题",
		Price:   23.5,
		Content: "商品详情",
	}

	addResponse := AddgoodsRes{}
	err2 := conn.Call("goods.AddGoods", addRequest, &addResponse)
	if err2 != nil {
		fmt.Println(err2)
	}
	// 4、 获取微服务server中的值
	fmt.Printf("%#v", addResponse)

	// 启动一个新的微服务调用

	var goodsData GetGoodsRes
	err3 := conn.Call("goods.GetGoods", GetGoodsReq{Id: 1}, &goodsData)
	if err3 != nil {
		fmt.Println(err3)
	}
	// 打印出getGoods的结果
	fmt.Printf("%#v", goodsData)
}
