package main

import "fmt"

func deferFun() {
	defer println("defer")
	fmt.Println("deferFul")
}
func main() {
	a := 1
	b := 2
	defer deferFun()
	defer fmt.Println("defer", b)
	fmt.Println(a)
	stackDefer()
}
func stackDefer() {
	name := "stack"
	fmt.Println(name)
	for _, v := range name {
		defer fmt.Print(string(v))
	}
}
