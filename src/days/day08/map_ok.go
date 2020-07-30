package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	fmt.Println(x, ok)
}
