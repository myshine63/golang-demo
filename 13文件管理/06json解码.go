package main

import (
	"encoding/json"
	"fmt"
)

type JsonStruct struct {
	Name     string            `json:"name"`
	Age      int               `json:"age"`
	Sex      bool              `json:"sex"`
	Language []string          `json:"language"`
	Addr     map[string]string `json:"addr"`
}

const jsonStr = `{
  "name": "tom",
  "age": 14,
  "sex": true,
  "language": [
    "go",
    "js",
    "node"
  ],
  "addr": {
    "phone": "12345",
    "address": "成都"
  }
}`

func main() {
	//changeJsonToMap()
	changeJsonToStruct()
}

// 将json转成map
func changeJsonToMap() {
	res := make(map[string]interface{}, 10)
	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		fmt.Println("转成map失败", err)
	}
	fmt.Printf("%+v\n", res)
	fmt.Println(res["addr"])
}

// 将json转成struct

func changeJsonToStruct() {
	res := new(JsonStruct)
	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		fmt.Println("解析jsons失败", err)
	}
	fmt.Printf("%+v\n", res)
}
