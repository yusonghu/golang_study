package main

import (
	error2 "days/day15/error"
	"fmt"
)

func main() {
	toError := error2.ToError()
	if toError != nil {
		fmt.Println(toError)
	}
	//error2.ToNet()
}
