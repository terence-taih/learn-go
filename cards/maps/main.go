package main

import "fmt"

func main() {
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

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

func removeFromMap(c map[int]string, key int) {
	delete(c, key)
}
