package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	jim.updateName("Joey")
	jim.print()
}

//if type of receiver is of pointer of person you are able to get by with
// not converting jim to a pointer ex. jimPointer := &jim <-- not needed
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}

// Value Types - Use pointers to update values with these types
// int
// float
// string
// bool
// structs

// Reference Types - No need to convert to pointers to update these types
// slices
// maps
// channels
// pointers
// functions
