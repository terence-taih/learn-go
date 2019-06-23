package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	map3()
}

func map1() {
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	colors["green"] = "#fffeee"
	printMap(colors)

	districts := make(map[int]string)
	districts[7] = "Seventh"
	districts[10] = "Tenth"

	removeFromMap(districts, 10)
	fmt.Println(districts)
}

func map3() {
	mymap := map[string]person{}
	mymap["thai"] = person{firstName: "Thai", lastName: "Nguyen"}
	fmt.Println(mymap)
}

func map2() {
	mymap := map[string][]string{
		"muoi": []string{"Playing", "Learning"},
	}
	mymap["thai"] = []string{"Football", "Music"}
	mymap["hang"] = []string{"Reading", "Cooking"}

	for p, fv := range mymap {
		fmt.Printf("The favorite things of %v is:\n", p)
		for i, v := range fv {
			fmt.Printf("%d: %v \t", i, v)
		}
		fmt.Print("\n")
	}
	delete(mymap, "thai")
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

func removeFromMap(c map[int]string, key int) {
	delete(c, key)
}
