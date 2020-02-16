package main

import "fmt"

func main() {
	aa := xstruct{
		A: 5,
	}
	aa.testPointer()
	fmt.Println(aa.A)

	ab := xstruct{
		A: 5,
	}
	ab.testValue()
	fmt.Println(ab.A)

}

type xstruct struct {
	A int
}

func (xs *xstruct) testPointer() {
	xs.A = 10
}

func (xs xstruct) testValue() {
	xs.A = 10
}

