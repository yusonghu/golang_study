package main

import "fmt"

func main() {
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	fmt.Println("原始 map == ", countryCapitalMap)
	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("删除元素后 map == ", countryCapitalMap)

}
