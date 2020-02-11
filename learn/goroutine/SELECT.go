package main

import "fmt"

//	在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。
//	Go内置了select关键字，可以同时响应多个通道的操作。
func main() {
	//可处理一个或多个channel的发送/接收操作。
	//如果多个case同时满足，select会随机选择一个。
	//对于没有case的select{}会一直等待，可用于阻塞main函数。
	ch := make(chan int, 20)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:

		}
	}
}
