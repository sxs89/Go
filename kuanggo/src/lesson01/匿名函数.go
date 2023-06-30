package main

import (
	"fmt"
)

func main() {
	f1()

	// 匿名函数
	f3 := func() {
		fmt.Println("我是f3函数")
	}
	f3()

	func() {
		fmt.Println("我是f3函数")
	}()

}

func f1() {
	fmt.Println("我是f1函数")
}

// 切片函数
func prinSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
