package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // custom type
}

func main() {
	ace := person{
		firstName: "portgas",
		lastName:  "d ace",
		contactInfo: contactInfo{
			email:   "ace@gmail.com",
			zipCode: 12345,
		},
	}

	ace.updateName("pino")
	ace.print()
}

// struct receiver
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// struct receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
