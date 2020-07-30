package main

import (
	"fmt"
)

func main() {
	name := "Hello World"
	for i := 0; i < len(name); i++ {
		fmt.Print(int(name[i]))
	}
	fmt.Printf("\n")
	for i := 0; i < len(name); i++ {
		fmt.Print(name[i])
	}

}
