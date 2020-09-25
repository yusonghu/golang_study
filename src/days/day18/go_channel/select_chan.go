package go_channel

import (
	"fmt"
	"time"
)

//	语法
//	select {
//    case communication clause  :
//       statement(s);
//    case communication clause  :
//       statement(s);
//    /* 你可以定义任意数量的 case */
//    default : /* 可选 */
//       statement(s);
//	}
//- 每个case都必须是一个通信
//
//- 所有channel表达式都会被求值
//
//- 所有被发送的表达式都会被求值
//
//- 如果有多个case都可以运行，select会随机公平地选出一个执行。其他不会执行。
//
//- 否则：
//
//  如果有default子句，则执行该语句。
//
//  如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

func chSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 100
	}()
	go func() {
		ch2 <- 200
	}()
	//	select 是 Go 中的一个控制结构。select 语句类似于 switch 语句。
	//	但是select会随机执行一个可运行的case。
	//	如果没有case可运行，它将阻塞，直到有case可运行。
	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2, ok := <-ch2:
		if ok {
			fmt.Println(val2)
		} else {
			fmt.Println("ch2 通道已经关闭")
		}
	case <-time.After(3 * time.Second):
		fmt.Println("超时")
	default:
		fmt.Println("执行了default。。")
	}

	time.Sleep(10 * time.Second)
}

func init() {
	chSelect()
}
