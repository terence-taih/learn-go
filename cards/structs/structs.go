package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//type deck []string
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo : antoher way, just use the struct type
}

func main() {
	testStruct()
}

func testStruct() {
	alex := person{firstName: "Alex",
		lastName: "Anderson",
		contact:  contactInfo{email: "test@gmail.com", zipCode: 1000}}
	fmt.Println(alex)
	var thai person
	thai.contact = contactInfo{email: "thainguyen@hoiio.com", zipCode: 100}
	(&thai).updateName("Terence")
	thai.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
