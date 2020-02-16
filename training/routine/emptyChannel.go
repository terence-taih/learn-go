package main

import (
	"fmt"
	"time"
)

func main() {
	testChannel := make(chan int)

	go func() {
		testChannel <- 1
	}()

	go func() {
		fmt.Println(<-testChannel)
	}()
	time.Sleep(1 * time.Second)
}
