package main

import "fmt"

type Controller struct {
	M int32
}

type Something interface {
	Get()
	Post()
}

func (c *Controller) Get() {
	fmt.Print("GET")
}

func (c *Controller) Post() {
	fmt.Print("POST")
}

type T struct {
	Controller
}

func (t *T) Get() {
	//new(test.Controller).Get()
	fmt.Print("T")
}
func main() {

	//	Duck Typing
	var something Something
	//	T的父类已经实现了接口的方法，所以T不需要实现接口的全部方法
	something = new(T)
	var t T
	t.M = 1
	something.Get()
}
