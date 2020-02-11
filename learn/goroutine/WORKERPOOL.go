package main

import (
	"fmt"
	"time"
)

//	将处理完的结果放入只写通道中，之后再从结果通道中取数据
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

//	goroutine 池
func main() {
	//	只读的通道
	jobs := make(chan int, 100)
	//	只写的通道
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	defer close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
