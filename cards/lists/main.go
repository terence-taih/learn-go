package main

import "fmt"

func main() {
	list1()
	fmt.Println(sum(1, 2, 3))

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(x...))

}

func list1() {
	fmt.Println("In list1")
	list := []int{1, 2, 3, 4, 5}
	for i, v := range list {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", list)

	y := []int{6, 7, 8}

	list = append(list, y...)
	list = append(list, 9)
	fmt.Println("x", list)

}

func sum(x ...int) int {
	fmt.Println("In sum")
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
