package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	p1 := Person{
		"张三",
		20,
		"男",
	}
	b, err := msgpack.Marshal(p1)
	if err != nil {
		fmt.Println("msgpack faild", err)
		return
	}
	var p2 Person
	err = msgpack.Unmarshal(b, &p2)
	if err != nil {
		os.Exit(-2)
	}
	fmt.Println(p2)
}
