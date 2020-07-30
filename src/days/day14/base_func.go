package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displaySalary() {
	fmt.Println(e.name)
	fmt.Println(e.salary)
	fmt.Println(e.currency)
}

type Blank struct {
	name string
}

//	定义相同得方法名需要赋予不同的结构体
func (b Blank) displaySalary() {
	b.name = "改变银行"
	fmt.Println(b)
}

func main() {
	emp := Employee{
		name:     "Sam",
		salary:   4000,
		currency: "$",
	}
	emp.displaySalary()
	blank := Blank{name: "天地银行"}
	blank.displaySalary()
}
