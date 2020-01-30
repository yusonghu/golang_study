package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
)

//定义一个结构体
type Student struct {
	Name    string
	Age     uint8
	Address string
}

func main() {
	// 序列化
	s1 := Student{
		Name:    "张三",
		Age:     19,
		Address: "中国",
	}
	var buffer bytes.Buffer
	//	创建编码器
	encoder := gob.NewEncoder(&buffer)
	//	编码
	err := encoder.Encode(s1)
	if err != nil {
		log.Fatal("Error", err)
	}
	bufferBytes := buffer.Bytes()
	fmt.Println("序列化：", bufferBytes)

	//	反序列化
	//	创建解码器
	decoder := gob.NewDecoder(bytes.NewReader(bufferBytes))
	var s2 Student
	err = decoder.Decode(&s2)
	if err != nil {
		log.Fatal("ERROR", err)
	}
	fmt.Println("反序列化：", s2)

}
