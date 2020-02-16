package main

import "fmt"

func panicFunc() {
	panic("Throw again")
}

func recover() interface{}{
	return "Test"
}

func main() {
	defer func() {
		r:= recover()
		if r != nil {
			fmt.Println("Recover successfully", r)
		}
	}()

	panicFunc()
}