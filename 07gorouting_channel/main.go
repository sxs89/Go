package main

import "fmt"

func main() {
	fmt.Println("sxs")

	// 1\创建一个管道
	ch := make(chan int, 3)

	// 2. 给管道里面存储数据
	ch <- 10
	ch <- 21
	ch <- 32

	// 3. 获取管道里的内容
	a := <-ch
	fmt.Println(a)
	<-ch // 这也是取值的一种方式

	c := <-ch
	fmt.Println(c)

	// 4.打印管道的长度 和 容量
	fmt.Printf("管道的值： %v 容量： %v, 长度%v\n", ch, cap(ch), len(ch))

	// 5. 管道类型
	var ch1 = make(chan int, 4)

	ch1 <- 34
	ch1 <- 54
	ch1 <- 34

	// 6、管道阻塞
	//ch6 := make(chan int, 1)
	//ch6 <- 34
	//ch6 <- 22 // all goroutines are asleep - deadlock!

	ch7 := make(chan int, 1)
	ch7 <- 66
	<-ch7
	ch7 <- 67
	d := <-ch7
	fmt.Println(d)

	var ch10 = make(chan int, 10)

	// 管道里赛值
	for i := 0; i < 10; i++ {
		ch10 <- i
	}

	// 循环数据
	for v := range ch10 {
		fmt.Println(v)
	}
}
