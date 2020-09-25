package go_channel

import (
	"fmt"
	"time"
)

//	time 包中的通道相关函数
//	Timer是一次性的时间触发事件，这点与Ticker不同，Ticker是按一定时间间隔持续触发时间事件。
//
//	Timer常见的创建方式：

func timerFun() {
	//	创建一个计时器,并定义时间
	timer := time.NewTimer(3 * time.Second)
	//	指针类型
	fmt.Printf("%T\n", timer)
	fmt.Println(time.Now())
	//此处在等待channel中的信号，执行此段代码时会阻塞3秒
	ch := timer.C
	fmt.Println(<-ch)
}

//	停止计时器
func stopTimer() {
	timer := time.NewTimer(3 * time.Second)
	//	新开启一个线程来处理触发后的事件
	go func() {
		//	等待触发信号
		//	timer停止后是无法获取到信号的
		<-timer.C
		fmt.Println("Timer 结束")
	}()
	time.Sleep(1 * time.Second)
	stop := timer.Stop()
	if stop {
		fmt.Println("Timer 停止了")
	}
	time.Sleep(5 * time.Second)
}

func afterTimer() {
	//它相当于NewTimer(d).C
	times := time.After(3 * time.Second)
	fmt.Printf("%T\n", times)
	fmt.Println(time.Now())
	t := <-times
	fmt.Println(t)
	//	golang中的Time格式化需要以”2006-01-02 15:04:05“的格式为准
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}

func init() {
	//timerFun()
	//stopTimer()
	//afterTimer()
}
