package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//map的基本使用
func baseMap() {
	baseMap := make(map[int]string, 8)
	baseMap[0] = "00"
	baseMap[1] = "01"
	baseMap[2] = "02"
	baseMap[3] = "03"
	fmt.Println(baseMap)
	fmt.Println(len(baseMap))
}

//判断是否是否存在
func isExist() {
	existMap := make(map[string]int)
	existMap["张三"] = 20
	existMap["李四"] = 80
	v, ok := existMap["王五"]
	if ok {
		fmt.Println("value == ", v)
	} else {
		fmt.Println("查无此人")
	}
}

//	遍历map
func rangeMap() {
	rangeMap := make(map[string]int)
	rangeMap["小明"] = 40
	rangeMap["小张"] = 90
	for k, v := range rangeMap {
		fmt.Println(k, v)
	}
}

//	按照指定顺序遍历map<原理就是将key进行排序，然后根据key取出元素>
func orderListMap() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

//map类型的切片
func sliceForMap() {
	mapSlice := make([]map[int]int, 10)
	fmt.Println(mapSlice)
	mapSlice[0] = make(map[int]int)
	mapSlice[0][1] = 1
	mapSlice[1] = make(map[int]int)
	//	必须指定到具体的哪个map才能打印出值
	fmt.Println(mapSlice[0])
}

//值为切片类型的map
func mapForSlice() {
	sliceMap := make(map[int][]string)
	fmt.Println(sliceMap)
	var value []string
	value = append(value, "你好", "我很好")
	sliceMap[0] = value
	fmt.Println(sliceMap)
}
func main() {
	//baseMap()
	//isExist()
	//rangeMap()
	//orderListMap()
	//sliceForMap()
	mapForSlice()
}
