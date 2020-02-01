package main

import "fmt"

//全局变量
var count int8 = 0

//函数定义(参数为a,b)
func sum(a int, b int) int {
	//局部变量
	var x int = 20
	return a + b + x
}

//可变参数(底层实际上是数组)
func sums(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

//	多返回值
func manyReturn(x, y int) (int, int) {
	return x, y
}

//返回值命名
func nameReturn(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//定义函数类型
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

//函数作为参数
func calc(x, y int, param func(int, int) int) int {
	return param(x, y)
}

//函数作为返回值
func funcReturn(way string) func(int, int) int {
	if way == "+" {
		return add
	} else if way == "-" {
		return sub
	} else {
		return nil
	}
}

//闭包(闭包=函数+引用环境)
func closure() func(int) int {
	var v int
	return func(k int) int {
		return v + k
	}
}

//	defer
//	defer会在方法返回后或程序结束后进行延迟处理，顺序是倒过来的
func deferFun() string {
	fmt.Println("start")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("end")
	return ""
}

//panic/recover错误处理
//panic可以在任何地方引发，但recover只有在defer调用的函数中有效
func errorOver() {
	fmt.Println("1")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover")
		}
	}()
	panic("error")
	fmt.Println("2")
}
func main() {
	//	函数的调用
	//sum(1, 2)
	//fmt.Println(sums(1, 2, 3, 4))
	//var c calculation
	//c = add
	//fmt.Println(c(1, 2))

	//deferFun()
	////匿名函数
	//func(x, y int) {
	//	fmt.Println(x + y)
	//}(1, 2)

	errorOver()
}
