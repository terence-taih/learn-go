package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://www.google.com")
	fmt.Println(*resp)

	fmt.Println("================ Content of html file: ================")
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
