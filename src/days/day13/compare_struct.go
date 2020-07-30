package main

import "fmt"

type name struct {
	firstName string
	lastName  string
}

func main() {
	//结构体是值类型，如果每个字段具有可比性，则是可比较的。
	//如果它们对应的字段相等，则认为两个结构体变量是相等的。
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}
	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	//	name5和name6均为对象类型，他们比较的是内存地址
	name5 := new(name)
	name5.firstName = "小明"
	name6 := new(name)
	name6.firstName = "小明"
	name7 := &name5
	fmt.Println(name5 == name6)
	fmt.Println(&name5, ":::::::", &name6)
	fmt.Println(*name7 == name5)
}
