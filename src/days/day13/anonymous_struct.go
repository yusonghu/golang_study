package main

import "fmt"

type Human struct {
	name           string
	age            int
	weight, height float64
}
type Student struct {
	Human //	这里可以作为继承关系的一个切人点,Human为匿名字段
	study string
	name  string
}

func main() {
	//	使用new创建的是对象类型，是无法进行比较的
	s := new(Student)
	//	若有相同名字的字段必须这样赋值
	s.Human.name = "张三丰"
	s.name = "张三"
	//	若无匿名字段中无这个属性则可以通过字段提升来赋值
	s.age = 222
	s.study = "我是三好学生"
	s.weight = 69.56
	s.height = 189
	fmt.Println(s)

	//	s4声明的是值类型变量
	var s4 Student
	s4.name = "小明"
	s4.study = "学渣"
	s4.Human = Human{
		name:   "明大仙",
		age:    12,
		weight: 43,
		height: 133,
	}
	fmt.Println(s4)
}
