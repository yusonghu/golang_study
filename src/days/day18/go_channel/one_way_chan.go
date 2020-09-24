package go_channel

import (
	"fmt"
)

//	只写通道
func writeData(ch chan<- int) {
	ch <- 100
}

//	只读通道
func readData(ch <-chan int) {
	fmt.Println("读到了", <-ch)
}

func init() {
	//ch := make(chan int)
	//go writeData(ch)
	//go readData(ch)
	//time.Sleep(1 * time.Second)
}
