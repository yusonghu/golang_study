package main

import (
	"fmt"
	"runtime"
	"time"
)

//	golang 调度器
//	Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。
//	默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）
func fa() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}
func fb() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

//Go语言中的操作系统线程和goroutine的关系：
//一个操作系统线程对应用户态多个goroutine。
//go程序可以同时使用多个操作系统线程。
//goroutine和OS线程是多对多的关系，即m:n。
func main() {
	//	设置并发时占用的cup核心数
	runtime.GOMAXPROCS(1)
	go fa()
	go fb()
	time.Sleep(2000)
}
