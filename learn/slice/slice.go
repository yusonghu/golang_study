package main

import "fmt"

//	切片<左闭右开>	切片的底层是数组
func main() {
	//arraySlice()
	//makeSlice()
	//copySlice()
	//eachSlice()
	appendSlice()
}

//	基于数组的切片
func arraySlice() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	b := a[0:4]
	fmt.Println("切片中的数据", b)
	fmt.Println("切片形式", a[:], a[:1], a[2:])
}

//使用make()函数构造切片
func makeSlice() {
	a := make([]int, 0, 20)
	fmt.Println(a)
	fmt.Println("长度", len(a))
	fmt.Println("容量", cap(a))
	//要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。
	fmt.Println(len(a) == 0)
}

//切片的赋值拷贝
//对一个切片的修改会影响另一个切片的内容
func copySlice() {
	s1 := make([]int, 3)
	s2 := s1
	s2[0] = 100
	fmt.Println(s2)
	fmt.Println(s1)
}

//	切片遍历
func eachSlice() {
	//切片的遍历方式和数组是一致的，支持索引遍历和for range遍历。
	s := []int{1, 3, 5}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
	for k, v := range s {
		fmt.Println(k, v)
	}
}

//	append()方法为切片添加元素
func appendSlice() {
	//append()函数将元素追加到切片的最后并返回该切片。
	//切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	a := []string{"追加1", "追加2"}
	citySlice = append(citySlice, a...)
	fmt.Println(citySlice)

}
