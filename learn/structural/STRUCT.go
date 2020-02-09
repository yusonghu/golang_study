package main

import "fmt"

//	普通的结构体
type City struct {
	cityName string
	no       string
}

//	组合结构体
type Person struct {
	name string
	age  int8
	sex  string
	city City
}

//基本实例化
func baseInstantiation() {
	var p Person
	p.name = "张三"
	p.age = 20
	p.city = City{
		cityName: "广州",
		no:       "003",
	}
	p.sex = "男"
	fmt.Println(p)
}

//匿名结构体
func anonymouStruct() {
	var p struct {
		name string
		city City
	}
	p.name = "李四"
	p.city = City{
		cityName: "益阳",
		no:       "004",
	}
	fmt.Println(p)
}

//创建指针类型结构体
//可以通过使用new关键字对结构体进行实例化
//创建的结构体是指针类型的
//Go语言中支持对结构体指针直接使用.来访问结构体的成员。
func pointStruct() {
	p := new(Person)
	fmt.Println(p)
}

//取结构体的地址实例化
//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
func pointAddressStruct() {
	p := &Person{}
	p.name = "王五"
	p.sex = "女"
	p.age = 20
	fmt.Println(p)
}

//使用键值对初始化
func keyValueStruct() {
	p := Person{
		name: "赵六",
		age:  0,
		sex:  "",
		city: City{
			cityName: "广州",
			no:       "005",
		},
	}
	fmt.Println(p)
}

//使用值的列表初始化
//若是简写的化顺序则不能改变
func simpleStruct() {
	p := &Person{
		"赵六",
		0,
		"男",
		City{
			"广州",
			"009",
		},
	}
	fmt.Println(p)
}

//构造函数
//Go语言的结构体没有构造函数,我们可以自己实现
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大
func NewPerson(name, sex string, age int8, city City) *Person {
	return &Person{
		name: name,
		age:  age,
		sex:  sex,
		city: city,
	}
}

func main() {
	//baseInstantiation()
	//anonymouStruct()
	//pointStruct()
	//pointAddressStruct()
	//simpleStruct()
	//NewPerson("钱七", "男", 20, City{"沅江", "023"})

}
