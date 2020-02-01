package main

import "fmt"

//	普通的指针
func generalPoint() {
	a := 10
	//	获取变量地址
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)
	//  指针取值（根据指针去内存取值）
	c := *b
	fmt.Printf("type of c %T\n", c)
	fmt.Printf("value of c %v\n", c)
}

//使用new的指针
//使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
func newPoint() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)
}

//	给指针赋值
func valForPoint() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)
}

//使用make分配内存
//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建
//它返回的类型就是这三个类型本身而不是他们的指针类型
func makeForPoint() {
	var b map[string]int
	b = make(map[string]int)
	b["test"] = 0
	fmt.Println(b)
}
func main() {
	//generalPoint()
	//newPoint()
	//valForPoint()
	makeForPoint()
}
