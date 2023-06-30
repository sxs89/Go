package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个channel，用户groutine之间传递数据
	ch := make(chan int)

	// 启动一个groutine
	go printNumbers(ch)
	// 4.打印管道的长度 和 容量
	fmt.Printf("管道的值： %v 容量： %v, 长度%v\n", ch, cap(ch), len(ch))
	//fmt.Println(ch)
	//return

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)

	time.Sleep(2 * time.Second)

}

func printNumbers(ch chan int) {
	// 从channel接收数据
	for num := range ch {
		fmt.Println("Received", num)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Done")
}

// 模拟用户登录
