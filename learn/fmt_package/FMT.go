package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("好熟悉的感觉啊")
	name := "golang"
	fmt.Printf("我是：%s\n", name)
	//	Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	//	Sprint系列函数会把传入的数据生成并返回一个字符串。
	getStr := fmt.Sprintln("我是传入的数据")
	fmt.Println(getStr)
	//Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	errStr := fmt.Errorf("我是一个错误")
	fmt.Println(errStr)
	//	获取输入值
	//	var(
	//		val string
	//		age int
	//	)
	//fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。
	//fmt.Scan(&val,&age)
	//fmt.Println("扫描结果",val,age)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容：")
	readString, _ := reader.ReadString('\n')
	fmt.Println("读到了", readString)
}
