package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	//content, err := ioutil.ReadFile(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Program error", err)
		os.Exit(1)
	}
	content := make([]byte, 10)
	isEndOfFile := false
	i := int64(0)
	for !isEndOfFile {
		_, err2 := file.ReadAt(content, i*10)
		i++
		if err2 != nil {
			isEndOfFile = true
		} else {
			fmt.Print(string(content))
		}
	}
	io.Copy(os.Stdout, file)
}
