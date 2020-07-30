package main

import "time"

//	结构体
type person struct {
	name string
	age  int
}

//	接口
type USB interface {
	link()
	unlink()
}

//	单独的类型
type myint int

//	函数类型
type add func(a, b int) float64

//	注意这里无法对一个非本地类型的类型添加方法

//type myDuration = time.Duration

//	可以改为类型使用
type myDuration time.Duration

func (m myDuration) EasySet() (s string) {
	//cannot define new methods on non-local type time.Duration
	return "yes"
}

func main() {
	var duration myDuration
	duration.EasySet()
}
