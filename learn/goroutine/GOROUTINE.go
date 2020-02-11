package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func hello() {
	defer wait.Done()
	fmt.Println("Hello Goroutine!")
}

//在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
//当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束
func main() {
	//启动单个goroutine(只有但等待组中的数字为0时才会退出等待)
	for i := 0; i < 10; i++ {
		//// 启动一个goroutine就登记+1
		wait.Add(1)
		go hello()
	}
	//	等待所有登记的goroutine都结束
	wait.Wait()
	fmt.Println("main goroutine done!")
}
