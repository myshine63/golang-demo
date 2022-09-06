package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", router())
	fmt.Println(err)
}

func tom(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello tom"))
}

func jerry(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello jerry"))
}

// 注册路由
func router() *http.ServeMux {
	// 路由转接器，可以注册多个路由和处理器
	mux := http.NewServeMux()
	mux.HandleFunc("/tom", tom)
	mux.HandleFunc("/jerry", jerry)
	return mux
}
