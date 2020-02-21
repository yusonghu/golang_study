package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.cnblogs.com/bananafish/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
		return
	}
	fmt.Println(string(all))
}
