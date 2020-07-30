package main

import "fmt"

func main() {
	var numbers []int
	printslice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printslice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printslice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printslice(numbers)

	/* 扩容 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	//	copy 后的两个切片是没有联系的
	copy(numbers1, numbers)
	printslice(numbers1)
}
func printslice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
