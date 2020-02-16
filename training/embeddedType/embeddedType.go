package main

import (
"log"
)

type User struct {
	Name  string
	Email string
}

type Person struct {
	Age  int
	Hometown string
}


type Admin struct {
	User
	Person
	Level string
}

func (u User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

func (u Person) Notify() error {
	log.Printf("Person: Sending %s<%s>\n",
		u.Age,
		u.Hometown)

	return nil	
}



type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func main() {
	admin := Admin{
		User: User{
			Name:  "john smith",
			Email: "john@email.com",
		},
		Level: "super",
	}

	SendNotification(admin)
}

