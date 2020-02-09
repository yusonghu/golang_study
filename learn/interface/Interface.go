package main

import "fmt"

//在Go语言中接口（interface）是一种类型，一种抽象的类型
//interface是一组method的集合，是duck-type programming的一种体现。
//接口（interface）是一种类型

//接口的定义
type person interface {
	say()
	run() string
}

type animal interface {
	sleep()
}

//接口嵌套,嵌套得到的接口的使用与普通接口一样
type Biology interface {
	person
	animal
}

//实现接口的条件
//一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。
//只要实现了接口中的所有方法，就实现了这个接口。
type xiaoming struct {
	name string
}

func (x *xiaoming) say() {
	fmt.Println(x.name, "hello world")
}

func (x *xiaoming) run() string {
	fmt.Println(x.name, "person running")
	return "我正在跑步"
}

//	一个对象可以实现多个接口
func (x xiaoming) sleep() {
	fmt.Println(x.name, "sleeping")
}

//	注意指针接收者只能是指针类型的对象才能使用
//	值接受这则可以是对象也可以是指针类型
func main() {
	x := new(xiaoming)
	x.name = "小明"
	x.say()
	x.run()
	x.sleep()
	//	空接口
	//	空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	//	空接口类型的变量可以存储任意类型的变量。
	var i interface{}
	i = "我是一个空接口"
	fmt.Println(i)
}
