package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//	socket 客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		// 读取用户输入
		input := bufio.NewReader(os.Stdin)
		str, err := input.ReadString('\n')
		if err != nil {
			println(err)
			return
		}
		if str == "Q\n" {
			fmt.Println("退出客户端")
			return
		}
		write, err := conn.Write([]byte(str))
		if err != nil {
			println(err)
			return
		}
		fmt.Println(write)
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
