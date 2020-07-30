package main

import "days/day14/oop/employee"

func main() {
	//	结构体和方法
	//e := employee.Employee{
	//	FirstName:   "Tony",
	//	LastName:    "teacher",
	//	TotalLeaves: 10,
	//	LeavesTaken: 18,
	//}
	//e.LeavesRemaining()	//	Tony teacher has -8 leaves remaining

	//使用New()代替构造函数(使用0值)
	//var e employee.Employee
	//e.LeavesRemaining()

	//使用构造函数
	e := employee.New("Sam", "todo", 23, 44)
	e.LeavesRemaining()
}
