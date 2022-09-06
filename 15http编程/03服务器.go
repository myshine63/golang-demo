package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handleFunc需要传入一个回调函数，函数类型为func(res http.ResponseWriter,req *http.Request)
	http.HandleFunc("/hello", myHandler)                 // 处理一般请求
	fileHandler := http.FileServer(http.Dir("./static")) // 会返回一个实现了Handler接口的对象
	// Handle需要传入实现了Handler接口的对象
	http.Handle("/tom", &Hello{})
	http.Handle("/static/", http.StripPrefix("/static/", fileHandler)) // 处理静态文件
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
func myHandler(res http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(res, "hello tom")
}

type Hello struct{}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello tom"))
}
