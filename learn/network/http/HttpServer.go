package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.PostFormValue("name"))
	fmt.Println(r.PostFormValue("age"))
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, "hello", string(all))
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		println("httpServer shutDown ", err)
		return
	}
}
