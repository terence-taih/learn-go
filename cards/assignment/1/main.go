package main

import "fmt"

const (
	x     = 40
	y int = 42
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func main() {
	fmt.Println(x, y)
	/*var a string
	var b string

	fmt.Scanf("%v", &a)
	fmt.Scanf("%s", &b)
	fmt.Println("a", a)
	fmt.Println("b", b)*/

	var literalString = `Here is a 
	very long 
	"String". you need to 
	separate by lines`
	fmt.Println(literalString)

	var c string
	fmt.Scanln(&c)
	fmt.Println("c", c)

	fmt.Println(South.String())

}

func (d Direction) String() string {
	return []string{"North", "East", "South", "West"}[d]
}
