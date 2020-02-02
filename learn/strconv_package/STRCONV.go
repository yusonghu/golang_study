package main

import (
	"fmt"
	"strconv"
)

//Atoi()函数用于将字符串类型的整数转换为int类型，函数签名如下。
func atoi() {
	s1 := "23"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		println(err)
		return
	}
	fmt.Println(i1)
}

//Itoa()函数用于将int类型数据转换为对应的字符串表示
func itoa() {
	i2 := 1024
	s := strconv.Itoa(i2)
	fmt.Println(s)
}

//Parse类函数用于转换字符串为给定类型的值
func parse() {
	strconv.ParseBool("false")
	strconv.ParseInt("23", 10, 32)
	strconv.ParseFloat("0.2233", 64)
	strconv.ParseUint("44", 10, 32)
}
func main() {
	//atoi()
	//itoa()
	parse()
}
