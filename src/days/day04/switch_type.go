package main

import (
	"fmt"
)

func main() {
	var x interface{} = func(int) float64 {
		return 23.3
	}
	switch x.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case func(int) float64:
		fmt.Println("func(int)")
	case bool, string:
		fmt.Println("bool,string")
	default:
		fmt.Println("未知类型")
	}
}
