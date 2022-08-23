package main

import (
	"encoding/json"
	"fmt"
)

type myJsonSchema struct {
	Name     string   `json:"name"`
	Language []string `json:"language"`
	Price    string   `json:"-"` // json解析时忽略这个
	Empty    string   `json:"empty,omitempty"`
}

func main() {
	changeMapToJson()
	changeStructToJson()
}

func changeMapToJson() {
	data := make(map[string]interface{}, 10)
	data["name"] = "computer"
	data["language"] = []string{"js", "go", "node"}
	// 用于json编码
	newJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json解析失败", err)
	} else {
		fmt.Println(string(newJson))
	}
	// 将json编码，方便格式化输出
	newFormatJson, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("json解析失败", err)
	} else {
		fmt.Println(string(newFormatJson))
	}
}

func changeStructToJson() {
	computer := myJsonSchema{
		Name:     "computer",
		Language: []string{"go", "js", "node"},
		//Empty:    "111",
	}
	computerJson, _ := json.MarshalIndent(computer, "", " ")
	fmt.Println(string(computerJson))
}
