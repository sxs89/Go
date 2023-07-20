package main

import (
	"fmt"
	"sync"
	"time"
)

//  go build -race main.go  表示可以查看是否存在共抢一块资源的问题

var count = 0
var wg = sync.WaitGroup{}
var mutex sync.Mutex

// 第二块
var m = make(map[int]int, 0)

func testMap(num int) {
	mutex.Lock()
	sum := 1
	for i := 0; i < num; i++ {
		sum *= i
	}
	m[num] = sum
	fmt.Println(m)
	time.Sleep(time.Millisecond * 1000)
	mutex.Unlock()
	wg.Done()
}

// 第一块   互斥锁
func test() {
	mutex.Lock() // 加锁
	count++
	fmt.Println("the count is :", count)
	time.Sleep(time.Millisecond * 1000)
	mutex.Unlock() // 解锁
	wg.Done()
}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		//go test()
		go testMap(i)
	}
	wg.Wait()
}
