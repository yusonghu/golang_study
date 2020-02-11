package main

import "fmt"

//	有缓冲的通道
func main() {
	// 创建一个容量为1的有缓冲区通道
	ch := make(chan int, 1)
	ch <- 10
	fmt.Println("发送成功")
}
