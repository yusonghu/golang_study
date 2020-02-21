package main

import (
	"fmt"
	"net"
)

func main() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		println(err)
		return
	}
	defer udp.Close()
	fmt.Println("服务端启动")
	for {
		var data []byte
		n, addr, err := udp.ReadFromUDP(data[:])
		if err != nil {
			println(err)
			return
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		// 发送数据
		_, err = udp.WriteToUDP(data[:n], addr)
		if err != nil {
			println(err)
			return
		}
	}
	fmt.Println("udp服务端关闭")
}
