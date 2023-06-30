package main

import (
	"fmt"
)

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Println("第一次的值为 %d\n", c)

	fmt.Println("----------------")

	var ab int = 100
	var bb int = 200
	var ret int

	ret = max(ab, bb)
	fmt.Println("最大值： %d\n", ret)

	var i int
	balance := [...]float32{100.2, 20.3, 50.0}

	// 切片
	//s := []int{1, 2, 3}
	//var numbers = make([]int, 1, 3, 4)

	for i = 0; i < 3; i++ {
		fmt.Printf("blance[%d] = %f\n", i, balance[i])
	}

	fmt.Printf("变量的地址为： %x\n", &i)

}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
