package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	b := 15
	var a int
	numbers := [6]int{1, 2, 3, 4}
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	for a < b {
		a++
		fmt.Println("a:", a)
	}
	for i, x := range numbers {
		fmt.Printf("第 %d 位的值是 %d", i, x)
	}
}
