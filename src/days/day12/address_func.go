package main

import "fmt"

func change(val *int) {
	*val = 20
}

//	数组不推荐使用地址修改值
func modify(sls *[3]int) {
	sls[0] = 80
}

//	数组推荐使用切片修改值
func modifyBySlice(sls []int) {
	sls[2] = 90
}

func main() {
	a := 99
	fmt.Println(a)
	change(&a)
	fmt.Println("changed", a)

	slice := [3]int{1, 2, 3}
	modify(&slice)
	fmt.Println("modify", slice)

	modifyBySlice(slice[:])
	fmt.Println("modifyBySlice", slice)
}
