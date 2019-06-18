package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	for l := range c {

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
