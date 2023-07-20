package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data  int
	mutex sync.Mutex
)

func processData() {
	//mutex.Lock()
	//defer mutex.Unlock()

	data++
	fmt.Println(data)
	time.Sleep(time.Millisecond * 10000)
}

func main() {
	/*fmt.Println("ceshi -------------start")

	str := "www.topgoer.com"

	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	fmt.Println(md5str1)

	fmt.Println("ceshi -------------end")*/

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processData()
			time.Sleep(time.Millisecond * 100)
		}()
	}

	wg.Wait()

	fmt.Println("Data:", data)

}
