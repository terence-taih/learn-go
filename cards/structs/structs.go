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
	contactInfo
	//contact   contactInfo
	// contactInfo : antoher way, just use the struct type
}

func main() {
	struct1()
	struct2()
}

func struct2() {
	// this one is called anonymous struct
	p := struct {
		firstName string
		lastName  string
	}{
		firstName: "Thai",
		lastName:  "Nguyen",
	}
	fmt.Println(p)
}

func struct1() {
	alex := person{firstName: "Alex",
		lastName:    "Anderson",
		contactInfo: contactInfo{email: "test@gmail.com", zipCode: 1000}}
	fmt.Println(alex)
	var thai person
	thai.contactInfo = contactInfo{email: "thainguyen@hoiio.com", zipCode: 100}
	(&thai).updateName("Terence")
	thai.print()

	// Notice this one: just use thai.email instead of thai.contactInfo.email
	fmt.Println("Special thing:", thai.email)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
