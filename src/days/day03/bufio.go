package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入一个字符串...")
	//	获取系统输入对象
	reader := bufio.NewReader(os.Stdin)

	readString, err := reader.ReadString('\n')
	if err != nil {
		println(err)
	}
	fmt.Println("度到的数据", readString)
}
