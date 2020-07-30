package main

import "fmt"

func main() {
	a := [...]float64{32.43, 432.1, 34.2, 1.3, 43.34}
	fmt.Println(len(a))
	fmt.Println(a)
	for i, x := range a {
		fmt.Println("index:", i, "  value:", x)
	}
}
