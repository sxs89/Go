package main

import (
	"code/service"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func test01(num int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("协程（%v）打印的第%v条数据\n", num, i)
	}
}

func main() {
	service.Login()
	//service.Sc()
	//service.Fors()
	//service.Shuzu()

	//go service.Say("world")
	//go service.Say("Hello")
	//service.Say("sxs")
	//service.Jiegouti()
	//service.Jiegouti2()
	//service.Zhizhen()
	//service.Qiepian()
	//service.Qiepian2()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go test01(i)
	}
	wg.Wait()
	println("主线程执行完毕。。。")

	// 指针地址值
	//service.Zhizhen

	//service.Maps()

}
