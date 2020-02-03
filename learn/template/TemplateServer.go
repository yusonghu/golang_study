package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//	这里name和age需要大写才能被外部访问到

type UserInfo struct {
	Name string
	Age  int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./learn/template/hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name: "张三",
		Age:  20,
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, user)

}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
