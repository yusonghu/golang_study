package main

import (
	"fmt"
)

//创建channel
var ch chan int

//	通道发送数据
func sendChannel(num int) {
	ch <- num
}

//	通道接收数据
func receiveChannel() {
	fmt.Println(<-ch)
}

//	通道关闭数据
func closeChannel() {
	close(ch)
}
func init() {
	ch = make(chan int)
}

//channel是一种类型，一种引用类型,默认值为nil
func main() {
	//	这种无缓冲的通道若只发不接收的化则会出现死锁情况，必须要接收通道中的内容
	go sendChannel(23)
	//go receiveChannel()
	//defer closeChannel()
}
