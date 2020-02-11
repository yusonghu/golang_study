package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。
var x int64
var wg sync.WaitGroup

//互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源
var lock sync.Mutex

//读写锁分为两种：读锁和写锁。
//当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
//当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
//读写锁非常适合读多写少的场景
var rwlock sync.RWMutex

//Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。
//如果要执行的函数f需要传递参数就需要搭配闭包来使用。
var oncelock sync.Once

//Go语言中内置的map不是并发安全的
var maplock = sync.Map{}

//	写共享数据
func wirte() {
	rwlock.Lock()
	x++
	//	假设读操作耗时200s
	time.Sleep(100 * time.Millisecond)
	rwlock.Unlock()
	wg.Done()
}

//	读共享数据
func read() {
	rwlock.Lock()
	// 假设读操作耗时1毫秒
	time.Sleep(1 * time.Millisecond)
	fmt.Println(x)
	rwlock.Unlock()
	wg.Done()
}

func add() {
	for i := 0; i < 5000; i++ {
		//	加锁
		lock.Lock()
		x = x + 1
		//	解锁
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	////	每有一个done后delia则会减一
	//wg.Add(2)
	//go add()
	//go add()
	////	只有当锁的delia为0时才会继续走下去
	//wg.Wait()
	//fmt.Println(x)

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go wirte()
	}
	fmt.Println(time.Now().Sub(start))
	for i := 0; i < 100; i++ {
		//	只会执行一次
		oncelock.Do(func() {
			fmt.Println("do it , just do it ")
		})
	}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			maplock.Store(key, n)
			value, _ := maplock.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
