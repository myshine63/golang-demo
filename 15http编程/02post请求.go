package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	myPost()
}
func myPost() {
	client := &http.Client{}
	body := io.NopCloser(strings.NewReader("name=tom&age=14"))
	req, _ := http.NewRequest("POST", "http://www.baidu.com", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	bodyBuf, _ := io.ReadAll(req.Body)
	fmt.Println(string(bodyBuf))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}
