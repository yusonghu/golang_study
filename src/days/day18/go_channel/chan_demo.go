package go_channel

import (
	"fmt"
	"time"
)

func createChannel() {
	var ch chan int
	if ch == nil {
		ch = make(chan int)
		fmt.Printf("数据类型是： %T", ch)
	}
	fmt.Printf("%T,%p\n", ch, ch)
	same(ch)
}

//	通道地址一致
//	通道本身就是指向地址
func same(ch chan int) {
	fmt.Printf("%T,%p\n", ch, ch)
}

func notifyData() {
	var ch1 chan bool
	fmt.Println(ch1)
	fmt.Printf("%T\n", ch1) //chan bool
	ch1 = make(chan bool)   //0xc0000a4000,是引用类型的数据
	fmt.Println(ch1)
	go func() {
		for i := 0; i < 1; i++ {
			fmt.Println("子goroutine中，i：", i)
		}
		ch1 <- true
	}()
	//	这里会阻塞程序
	data := <-ch1
	fmt.Println("data-->", data)
	fmt.Println("main。。over。。。。")
}

//	死锁
func deadlock() {
	//如果Goroutine正在等待从通道接收数据，那么另一些Goroutine将会在该通道上写入数据，否则程序将会死锁。
	ch := make(chan int)
	go func() {
		ch <- 5
	}()
	result := <-ch
	fmt.Println(result)
	time.Sleep(2000)
}

func closeChan() {
	ch := make(chan int)
	go sendData(ch)
	for true {
		time.Sleep(1 * time.Second)
		v, ok := <-ch
		if !ok {
			fmt.Println("读取结束")
			break
		}
		fmt.Println("读到数据：", v)
	}
	fmt.Println("OVER。")
}

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		//	向通道中发送数据
		time.Sleep(200 * time.Millisecond)
		ch <- i
	}
	//	关闭通道
	close(ch)
}

//	通道上的范围循环
func chForRange() {
	// for循环的for range形式可用于从通道接收值，直到它关闭为止。
	ch := make(chan int)
	go sendData(ch)
	for i := range ch {
		fmt.Println("读取数据:", i)
	}
	fmt.Println("数据已经读完，over")
}

func init() {
	//createChannel()
	//notifyData()
	//deadlock()
	//closeChan()
	//chForRange()
}
