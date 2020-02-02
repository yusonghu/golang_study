package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ioutilCopyFile() {
	file, err := ioutil.ReadFile("./logTest.log")
	if err != nil {
		println(err)
		return
	}
	err = ioutil.WriteFile("log.log", file, 0777)
	if err != nil {
		println(err)
		return
	}
}

func bufioCopyFile() {
	open, err := os.Open("./logTest.log")
	if err != nil {
		println(err)
		return
	}
	defer open.Close()
	file, err := os.OpenFile("bufiolog.log", os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		println(err)
		return
	}
	var writer = bufio.NewWriter(file)
	//	边读边写
	reader := bufio.NewReader(open)
	for {
		line, err := reader.ReadString('\n')
		writer.WriteString(line)
		if err == io.EOF {
			fmt.Println("读完了")
			break
		}
		if err != nil {
			println(err)
			return
		}
	}
	writer.Flush()
}
func main() {
	//ioutilCopyFile()
	bufioCopyFile()
}
