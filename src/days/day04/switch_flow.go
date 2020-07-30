package main

import "fmt"

func main() {
	var grade string = "B"
	var marks int = 80
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
		//	fallthrough将继续往下跑
		fallthrough
	case 70, 60:
		grade = "C"
	default:
		grade = "D"
	}
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("及格")
	case "D":
		fmt.Println("不及格")
	}
	fmt.Println("您的等级为:", grade)
}
