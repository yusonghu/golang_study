package main

import "fmt"

func main() {
	var countryMap map[string]string
	//	map 的 == 运算用来判断是否为空
	fmt.Println(countryMap == nil)
	countryMap = make(map[string]string)
	fmt.Println(countryMap)
	countryMap["china"] = "beijing"
	countryMap["jepan"] = "beijing"
	for index, country := range countryMap {
		fmt.Println(index, "===", country)
	}
	v, ok := countryMap["english"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}

}
