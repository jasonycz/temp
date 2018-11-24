package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //定义一个同步等待的组

type HH []int

var H HH

func task(i int, h *HH) {
	fmt.Println("task...", i)
	*h = append(*h, i)
	//耗时操作任务，网络请求，读取文件
	time.Sleep(2 * time.Second)
	wg.Done() //减去一个计数
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1) //添加一个计数
		go task(i, &H)
	}
	wg.Wait() //阻塞直到所有任务完成
	fmt.Println("over")
	fmt.Printf("%+v", H)
}
