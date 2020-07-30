package main

import "fmt"

func main() {
	a := [5]int{53, 43, 2, 421, 3}
	//	左闭右开
	var b []int = a[1:4]
	fmt.Println(b)
}
