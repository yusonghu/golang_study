package main

import "fmt"

type name int8

type first struct {
	a int
	b bool
	name
}

func main() {
	var a int = 20
	var ip *int = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)
	fmt.Printf("ip 变量的存储地址: %x\n", ip)
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	f := new(first)
	f.a = 1
	f.name = 12
	fmt.Println(f.a, f.b, f.name)

	second := first{1, false, 2}
	x := &second
	fmt.Println(*x, x)

}
