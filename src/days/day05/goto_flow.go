package main

import "fmt"

func main() {
	var a int = 10

	for a < 20 {
		if a == 15 {
			//	跳到LOOP,直接退出循环,可使用goto统一处理错误
			goto END
		}
		fmt.Println(a)
		a++
	}
END:
	fmt.Println("2333")
}
