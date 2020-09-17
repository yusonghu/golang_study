package go_resource

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//	共享资源
func awwResourcce() {
	a := 1
	go func() {
		//	共享a的数据
		a = 2
		fmt.Println("子goroutine。。", a)
	}()
	a = 3
	fmt.Println("main goroutine。。", a)
}

//	定义一个全局变量
var ticket = 10

//	创建一个同步等待组,保证所有协程全部结束后主协程再结束
var wg sync.WaitGroup

// 创建互斥锁
//每个资源都对应于一个可称为 “互斥锁” 的标记，这个标记用来保证在任意时刻，只能有一个协程（线程）访问该资源。其它的协程只能等待。
var matex sync.Mutex

func awwResourceWarning() {
	//设置WaitGroup的计数的值
	wg.Add(4)
	go saleTickets("窗口1")
	go saleTickets("窗口2")
	go saleTickets("窗口3")
	go saleTickets("窗口4")
	//进入阻塞状态。一直到WaitGroup的计数器为零。才能解除阻塞。
	wg.Wait()
	//time.Sleep(5 * time.Second)
}

//	这里没有任何保护措施，售出的票总和会比总票数还多
//	为解决这种问题我们需要加锁

//	PS:不要以共享内存的方式去通信，而要以通信的方式去共享内存
func saleTickets(name string) {
	nano := time.Now().UnixNano()
	//	若nano不一样则每次生成出来的随机数和上次生成的随机数不一样
	rand.Seed(nano)
	//将WaitGroup中的计数值减1，若WaitGroup中的计数值为负数则会抛出异常
	defer wg.Done()
	for {
		//	加锁，不能重复加
		matex.Lock()
		if ticket <= 0 {
			matex.Unlock()
			fmt.Println("窗口<" + name + ">没票了")
			break
		} else {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket)
			ticket--
		}
		//	释放锁，不能重复加
		matex.Unlock()
	}
}

// 读写锁
//1. 同时只能有一个 goroutine 能够获得写锁定。
//2. 同时可以有任意多个 gorouinte 获得读锁定。
//3. 同时只能存在写锁定或读锁定（读和写互斥）。

//	基本遵循两大原则：
//1、可以随便读，多个goroutine同时读。
//
//2、写的时候，啥也不能干。不能读也不能写。
var rwMutex *sync.RWMutex

func rwLock() {
	//	初始化读写锁
	rwMutex = new(sync.RWMutex)
	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)
	wg.Wait()
}

func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始写：write start 。。。")
	//	写操作上锁
	rwMutex.Lock()
	fmt.Println(i, "正在写：writing。。。")
	time.Sleep(2 * time.Second)
	//	写操作解锁
	rwMutex.Unlock()
	fmt.Println(i, "写结束：write over 。。。")
}

func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始读：read start。。。")
	//	读取操作上锁
	rwMutex.RLock()
	fmt.Println(i, "正在读：reading。。。")
	time.Sleep(3 * time.Second)
	//读解锁
	rwMutex.RUnlock()
	fmt.Println(i, "读结束：read over。。。")
}

func init() {
	//awwResourceWarning()
	rwLock()
}
