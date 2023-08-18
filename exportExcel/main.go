package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"sync"
	"time"
)

type Data struct {
	Id    int
	Name  string
	Score int
}

func main() {

	data := []Data{
		{1, "Alice1", 85},
		{2, "Zone", 66},
		{3, "sxs", 55},
	}

	var wg sync.WaitGroup
	var mu sync.Mutex // 互斥锁

	// 创建一个Excel文件

	file := excelize.NewFile()

	for _, d := range data {
		wg.Add(1)
		go func(data Data) { // 开启多协程任务处理批量导出数据
			defer wg.Done()

			// 在写入Excel之前锁定互斥对象
			mu.Lock()
			defer mu.Unlock()

			//fmt.Println("Data=>", data)
			time.Sleep(time.Millisecond * 100)
			file.SetCellValue("Sheet1", fmt.Sprintf("A%d", data.Id), data.Id)
			file.SetCellValue("Sheet1", fmt.Sprintf("B%d", data.Id), data.Name)
			file.SetCellValue("Sheet1", fmt.Sprintf("C%d", data.Id), data.Score)
		}(d)
	}

	wg.Wait()

	err := file.SaveAs("data.xlsx")
	if err != nil {
		fmt.Println("导出文件报错", err)
	} else {
		fmt.Println("多协程导出数据成功！")
	}
}
