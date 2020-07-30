package conf

import "fmt"

type conf struct {
	name  string
	value string
}

func GetInfo() {
	fmt.Println("conf")
}
func init() {
	fmt.Println("conf init ... ")
}
