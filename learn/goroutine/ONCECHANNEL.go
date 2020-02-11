package main

import "fmt"

//	单向通道
//	限制通道在函数中只能发送或只能接收(单向通道也是通道的类型)
//	<- chan int 只读类型的单向通道
//	chan<- int 只写类型的单向通道

//	只写通道
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

//	读写通道
func squarer(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i
	}
	close(out)
}

//	只读通道
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch1, ch2)
	printer(ch2)
}
