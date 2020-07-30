package main

import "fmt"

func main() {
	var a int = 10
	if a <= 20 {
		fmt.Println("小于等于20")
	} else {
		fmt.Println("大于20")
	}
	if num := 10; num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

}
