package main

import (
	"fmt"
	"sync"
	"time"
)

/*func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
}

func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}*/

/*func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}*/

func calculateAverage(num float64, resultCh chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()

	// 模拟计算延迟
	// 在实际情况下，可以在这里执行实际的计算操作
	// 这里仅模拟计算延迟
	for i := 0; i < 100000000; i++ {

	}

	time.Sleep(1000 * time.Millisecond)

	// 计算平均值
	average := num / 2.0
	resultCh <- average
	fmt.Println("num ====>", num)
}

func main() {
	//ch := make(chan string, 1)
	//send(ch, "hello world")
	//read(ch)

	/*go say("php")
	say("java")*/

	//numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers := []float64{1, 2, 3, 4}

	// 创建一个无缓存的通道，用于传递计算结果
	resultCh := make(chan float64)
	// 创建一个 WaitGroup，用于等待所有计算完成
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go calculateAverage(num, resultCh, &wg)
	}

	// 启动一个 goroutine 等待所有计算完成
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 从通道中读取结果并计算平均值
	var sum float64
	count := 0
	// resultCh 表示通道中获取每个值的平均值
	for avg := range resultCh {
		sum += avg
		count++
	}
	// 计算并输出平均值
	if count > 0 {
		totalAverage := sum / float64(count)
		fmt.Println("Total Average: %.2f\\n", totalAverage)
		fmt.Println("count的值为：", count)
	} else {
		fmt.Println("No data to calculate average.")
	}

}
