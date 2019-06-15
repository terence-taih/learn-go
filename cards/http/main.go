package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://www.google.com")
	fmt.Println(*resp)
}
