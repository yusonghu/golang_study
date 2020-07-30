package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 200
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	fmt.Println(add(&a))
	ret = max(a, b)
	fmt.Printf("最大值是 : %d\n", ret)

}
func add(a *int) int { // 请注意，
	*a = *a + 1 // 修改了a的值
	return *a   // 返回新值
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
