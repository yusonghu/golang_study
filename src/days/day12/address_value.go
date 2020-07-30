package main

import "fmt"

func main() {
	a := 255
	b := &a
	fmt.Println(*b)
	//	GO不支持指针算法，所以需要获取到指针的值进行计算
	*b++
	fmt.Println(*b)
	fmt.Println(a)
	array := []int{10, 100, 200}
	var ptr [3]*int
	var i int

	for i = 0; i < len(array); i++ {
		ptr[i] = &array[i]
	}
	for _, v := range ptr {
		fmt.Println(v, *v)
	}
}
