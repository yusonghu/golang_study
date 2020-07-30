package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i > 8 {
			//	直接跳出循环
			break
		}
		if i == 3 {
			//	跳出循环
			continue
		}
		fmt.Println(i)
	}
}
