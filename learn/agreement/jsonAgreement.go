package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//	若需要转json的化参数首字母需要大写
type jsonData struct {
	Id  int    `json:"id"`
	Val string `json:"value"`
}

func main() {
	// 1. 使用 json.Marshal 编码
	j1 := jsonData{
		Id:  1,
		Val: "01",
	}
	m, _ := json.Marshal(&j1)
	fmt.Println(string(m))
	// 2. 使用 json.Unmarshal 解码
	var j2 jsonData
	json.Unmarshal(m, &j2)
	fmt.Println(j2)
	// 3. 使用 json.NewEncoder 编码
	j3 := jsonData{
		Id:  2,
		Val: "02",
	}
	// 编码结果暂存到 buffer
	bytes := new(bytes.Buffer)
	json.NewEncoder(bytes).Encode(&j3)
	fmt.Println(string(bytes.Bytes()))
	// 4. 使用 json.NewDecoder 解码
	var j4 jsonData
	json.NewDecoder(bytes).Decode(&j4)
	fmt.Println(j4)
}
