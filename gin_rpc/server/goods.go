package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Goods struct {
}

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

func (this Goods) AddGoods(req AddgoodsReq, res *AddgoodsRes) error {
	// 业务逻辑处理
	fmt.Printf("%#v", req)
	// 返回增加的结果

	// 2. 返回结果
	*res = AddgoodsRes{
		Success: true,
		Message: "执行成功",
	}
	return nil
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

func (this Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	// 1、执行增加模拟
	fmt.Printf("%#v", req)
	// 2.返回的数据结果
	*res = GetGoodsRes{
		Id:      10,
		Title:   "PHP",
		Price:   19.9,
		Content: "PHP修炼之路",
	}
	return nil
}

// ---------------------------------------------

func main() {
	//	1. 注册RPC服务
	err1 := rpc.RegisterName("goods", new(Goods))
	if err1 != nil {
		fmt.Println(err1)
	}

	// 2、监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8020")
	if err2 != nil {
		fmt.Println(err2)
	}
	// 3.关闭监听
	defer listener.Close()

	for {
		// 4.监听客户端的连接
		fmt.Println("准备建立连接-------------")
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		rpc.ServeConn(conn)
	}
}
