package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//	打开和关闭文件
func openAndCloseFile() {
	//只读的方式打开文件
	open, err := os.Open("./logTest.log")
	//	关闭文件
	defer open.Close()
	if err != nil {
		fmt.Println(err)
	}
}

//读取文件(基本读取)
func readFile() {
	file, err := os.Open("./logTest.log")
	defer file.Close()
	if err != nil {
		println(err)
		return
	}
	//	使用read方法读取数据
	var temp = make([]byte, 1024)
	read, err := file.Read(temp)
	if err != nil {
		println(err)
		return
	}
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	fmt.Printf("读取的字节数%v\n", read)
	fmt.Println(string(temp[:read]))
}

//读取文件（循环读取）
func rangeRead() {
	file, err := os.Open("./logTest.log")
	defer file.Close()
	if err != nil {
		print("未找到该文件")
		return
	}
	var content []byte
	temp := make([]byte, 1024)
	for {
		read, err := file.Read(temp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			println(err)
			return
		}
		content = append(content, temp[:read]...)
	}
}

//bufio读取文件
func bufioRead() {
	file, err := os.Open("./logTest.log")
	defer file.Close()
	if err != nil {
		print("未找到该文件")
		return
	}
	reader := bufio.NewReader(file)
	for {
		//	一行一行读取
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			println(err)
			return
		}
		fmt.Println(line)
	}
}

//ioutil读取整个文件
func ioutilRead() {
	file, err := ioutil.ReadFile("./logTest.log")
	if err != nil {
		println(err)
		return
	}
	fmt.Println(string(file))
}

//	Write和WriteString
func writeAndWriteString() {
	//os.O_WRONLY	只写
	//os.O_CREATE	创建文件
	//os.O_RDONLY	只读
	//os.O_RDWR		读写
	//os.O_TRUNC	清空
	//os.O_APPEND	追加
	//文件权限,采用的是unix文件权限的数字表示方式(采用8进制)
	file, err := os.OpenFile("wirte.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		println(err)
		return
	}
	defer file.Close()
	str := "hello world"
	file.Write([]byte(str))
	file.WriteString(str)
}

//bufio创建文件
func bufioWriter() {
	file, err := os.OpenFile("bufioWirte.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "hello go"
	//	这个会将数据写入缓存中
	writer.Write([]byte(str))
	writer.WriteString(str)
	//将缓存中的数据写入文件中
	writer.Flush()
}

//ioutil写入文件
func ioutilWriter() {
	str := "hello ioutil"
	err := ioutil.WriteFile("ioutilWrite.txt", []byte(str), 0777)
	if err != nil {
		println(err)
		return
	}
}

func main() {
	//openAndCloseFile()
	//readFile()
	//rangeRead()
	//bufioRead()
	//ioutilRead()
	//writeAndWriteString()
	//bufioWriter()
	ioutilWriter()
}
