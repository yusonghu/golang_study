package go_runtime

import (
	"fmt"
	"runtime"
)

func RuntimeFunc() {
	//	runtime 包中常用的函数
	//	CPU 核数
	fmt.Println(runtime.NumCPU())
	//	设置最大的可同时使用的 `CPU` 核数
	//	go1.8后，默认让程序运行在多个核上,可以不用设置了
	//	go1.8前，还是要设置一下，可以更高效的利益CPU
	runtime.GOMAXPROCS(4)
	//让当前线程让出 `cpu` 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("groutine...")
		}
	}()
	runtime.Gosched()
	//返回正在执行和排队的任务总数
	//注意：垃圾回收所在Groutine的状态也处于这个范围内的话，也会被纳入该计数器。
	fmt.Println(runtime.NumGoroutine())
	//目标操作系统
	//很多时候，我们会根据平台的不同实现不同的操作，就而已用GOOS了
	fmt.Println(runtime.GOOS)
	//会让运行时系统进行一次强制性的垃圾收集
	runtime.GC()
	//获取goroot目录
	fmt.Println(runtime.GOROOT())
	defer println("defer in goroutine")
	//退出当前 `goroutine`(但是`defer`语句会照常执行)
	runtime.Goexit()
}
