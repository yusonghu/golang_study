package main

import (
	"fmt"
	"net"
)

func main() {
	udpClient, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		println(err)
		return
	}
	defer udpClient.Close()
	sendData := []byte("hello udp server")
	_, err = udpClient.Write(sendData)
	if err != nil {
		println(err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := udpClient.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
