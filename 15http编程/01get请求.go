package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//myDo()
	//setReq()
	myGet()
}

// 通过do方法发起请求
func myDo() {
	// 创建一个客户端
	client := &http.Client{}
	// 创建一个请求对象
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("创建请求对象失败")
	}
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发起请求失败", err)
		return
	}
	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println("Status:", resp.Status)
	fmt.Printf("Request:%+v\n", resp.Request)
	bodyBuf, _ := io.ReadAll(resp.Body)
	fmt.Println("body:", string(bodyBuf))
}

// 设置请求参数
func setReq() {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	// 创建cookie对象
	cookie := &http.Cookie{
		Name:  "name",
		Value: "tom",
	}
	// 添加cookie
	req.AddCookie(cookie)
	// 设置Header
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}

// 通过简化的get发送请求
func myGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}
