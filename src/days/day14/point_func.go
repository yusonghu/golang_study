package main

import "fmt"

type Rectangle struct {
	width, height int
}

//	指针类型会将以前的数据直接修改
//	若无这个指针是无法修改值的
func (r *Rectangle) setVal() {
	r.height = 20
	r.width = 44
}

func main() {
	rectangle := Rectangle{}
	rectangle.setVal()
	fmt.Println(rectangle)
}
