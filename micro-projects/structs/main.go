package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Go adheres to "pass by value" which means when a value is updated, a new struct, with the data of the old struct copied over and the new updated value, is created at an adjacent memory address.
// To resolve this, we can use pointers, which "point to a memory address".
// *person is a type description - it means we're working with a pointer to a type person.
// *pointerToPerosn is an operator, which means we want to manipulate the value the pointer is referencing.
// &variable --> give me the memory address the value is pointing at.
// *variable --> give me the value the memory address is pointing at.
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Turn address into value with *address
// Turn value into address with &value

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contact: contactInfo{
			"jim@gmail.com",
			123,
		},
	}

	// Point to and give me access to the memory address of jim.
	jim.updateName("jimmy")
	jim.print()
}
