package main

import "fmt"

var test int

func main() {
	testStruct()
}

func quiz1() {
	ints := []int{}
	for i := 0; i < 10; i++ {
		ints = append(ints, i)
	}

	for j := range ints {
		if j%2 == 0 {
			fmt.Println(j, " is even")
		} else {
			fmt.Println(j, " is old")
		}
	}

}
