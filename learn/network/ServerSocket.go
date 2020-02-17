package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buffer [1024]byte
	for {
		//	读取数据
		n, err := reader.Read(buffer[:])
		if err != nil {
			println(err)
			break
		}
		recvStr := buffer[:n]
		result := "服务端收到数据" + string(recvStr)
		fmt.Println("收到客户端发来的数据:", string(recvStr))
		conn.Write([]byte(result))
	}
}

//	网络服务端
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("服务端开始监听...")
		accept, err := listen.Accept()
		fmt.Println("服务端监听到了连接")
		if err != nil {
			panic(err)
		}
		go process(accept)
	}

}
