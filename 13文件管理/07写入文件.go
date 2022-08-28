package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writeFile("hello.txt")
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	bufWrite := bufio.NewWriter(file)
	defer bufWrite.Flush()
	for i := 0; i < 10; i++ {
		_, _ = fmt.Fprintln(bufWrite, i)
	}
}
