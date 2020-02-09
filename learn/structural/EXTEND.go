package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	name string
}

type Dog struct {
	call    string
	*Animal //	通过匿名嵌套结构体实现继承
}

func (d *Dog) wang() {
	fmt.Println(d.Animal.name, "汪汪")
}

//	动物这个结构体所具有的方法
func (a *Animal) move() {
	fmt.Println(a.name, "会动")
}

//结构体标签（Tag）
type Student struct {
	ID   int    `json:"id"`
	name string //	私有的不能被json包访问
	Age  int
}

//	json 序列化
func (stu *Student) JSONSerialize() (err error) {
	marshal, err := json.Marshal(stu)
	if err != nil {
		return err
	}
	fmt.Println(string(marshal))
	return nil
}

//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
func main() {
	//dog := &Dog{
	//	call:   "頑張って",
	//	Animal: &Animal{"依依"},
	//}
	//dog.wang()
	//dog.move()

	stu := new(Student)
	stu.name = "辣个男人"
	stu.Age = 30
	stu.ID = 400
	stu.JSONSerialize()
}
