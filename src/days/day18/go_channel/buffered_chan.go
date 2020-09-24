package go_channel

import (
	"fmt"
	"log"
	"time"
)

//	缓冲通道
func bufferedChan() {
	//缓冲通道，缓冲是5
	ch1 := make(chan int, 10)
	fmt.Println("len", len(ch1), "cap", cap(ch1))
	ch1 <- 2
	fmt.Println("len", len(ch1), "cap", cap(ch1))
	log.Println("-----------")
	go sendData(ch1)
	for {
		time.Sleep(1 * time.Second)
		v, ok := <-ch1
		if !ok {
			fmt.Println("读完了")
			break
		}
		fmt.Println("读到了", v)
	}
}

func init() {
	//bufferedChan()
}
