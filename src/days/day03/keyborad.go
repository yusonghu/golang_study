package main

import "fmt"

func main() {
	var x int
	var y float64
	fmt.Println("请输入一个整数和一个浮点数")
	//读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
	fmt.Scanln(&x, &y)
	fmt.Printf("x的数值：%d，y的数值：%f\n", x, y)
	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d,y:%f\n", x, y)
}