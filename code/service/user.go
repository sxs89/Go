package service

import (
	"fmt"
	"time"
)

func Login() {
	fmt.Println("hello world")
	fmt.Println("你好，我的世界!!!")
}

// 和函数实现
func Sc() {
	var a int = 21
	var b int = 30
	var c int
	c = a + b
	fmt.Println("两个值相加的的结果是", c)
}

// 循环函数实现
func Fors() {
	for true {
		fmt.Printf("这是无限循环 \n")
	}
}

// 数组的实现
func Shuzu() {
	var n [10]int
	var i, j int

	// 为数组添加数据
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

// go 的并发实现
func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 创建一个结构体
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// 结构体测试一
func Jiegouti() {
	fmt.Println(Books{"Go语言", "www.runoob.com", "Go 语言教程", 4564})

	fmt.Println(Books{title: "PHP", subject: "php 是世界上最好的语言"})
}

// 结构体测试二
func Jiegouti2() {
	var Book1 Books
	var Book2 Books

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}

// 指针
func Zhizhen() {
	var a int = 10
	fmt.Printf("变量的地址：%x\n", &a)
}

// 切片
func Qiepian() {
	var numbers = make([]int, 2, 3)
	PrintSlice(numbers)
}

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%d\n", len(x), cap(x), x)
}

func Qiepian2() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("原始切片数据 ==", numbers)
	fmt.Println("=================")
	numbers = append(numbers, 22)
	fmt.Println("追加后的原始切片数据 ==", numbers)

	for key, value := range numbers {
		fmt.Printf("key==%d, value=%v\n", key, value)
	}

}

// 指针
func zhizhen() {
	var a int = 10
	fmt.Printf("变量的地址值%x", &a)
}

// Map
func Maps() {
	var countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	fmt.Println("-------删除后---------")
	delete(countryCapitalMap, "Japan")
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

}
