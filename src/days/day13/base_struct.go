package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	var wuniangaokao Books
	var sannianmoni Books

	wuniangaokao.title = "五年高考"
	wuniangaokao.author = "男的女的折磨"
	wuniangaokao.subject = "高考有这一本就可以了"
	wuniangaokao.bookId = 393012
	printBook(&wuniangaokao)
	sannianmoni.title = "三年模拟"
	sannianmoni.author = "男的女的折磨"
	sannianmoni.subject = "我是魔鬼"
	sannianmoni.bookId = 29438
	printBook(&sannianmoni)
	fmt.Println(Home{})
}

type Home struct {
}

//	对象的字符串输出格式
func (h Home) String() string {
	return "这个方法在打印输出这个对象的时候会自动被调用"
}

//传入指针类型，使用同一个地址，否则会拷贝这个结构体
func printBook(book *Books) {
	fmt.Println(&book.title)
	fmt.Println(book.author)
	fmt.Println(book.subject)
	fmt.Println(book.bookId)
}
