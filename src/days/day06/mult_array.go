package main

import "fmt"

//	二位数组
func main() {
	var a [3][4]int = [3][4]int{
		{3, 1, 2, 3},
		{3, 4, 4, 5},
		{2, 4, 4, 3},
	}
	for _, av := range a {
		fmt.Println("av:", av)
		for _, bv := range av {
			fmt.Println("bv:", bv)
		}
	}
}
