package main

import (
	"fmt"
	"net/http"
	"time"
)

// 中间件其实就是对handlerFunc的封装，也可以说是高阶函数，它接收一个handler,并返回一个处理后的handler
func LogTime(handler http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Println(start)
	})
}

func main() {

}
