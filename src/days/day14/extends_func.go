package main

import "fmt"

type Animal struct {
	name string
}

type cat struct {
	Animal
	master string
}

func (a Animal) eat() {
	fmt.Println("动物会吃东西...")
}

func (c cat) say() {
	fmt.Println("猫会叫....喵喵喵!")
}

//	重写方法
func (c cat) eat() {
	fmt.Println("猫吃鱼...")
}

func main() {
	c := cat{
		Animal: Animal{name: "猫"},
		master: "张三",
	}
	c.eat()
	c.say()
}
