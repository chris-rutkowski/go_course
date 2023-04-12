package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // short for contactInfo contactInfo - field and type the same name
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{lastName: "Anderson", firstName: "Alex"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person // currently empty string first and last names
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex) // field names and their values e.g. `{firstName: Alex lastName:Anderson}``

	jim := person{
		firstName: "Jim",
		lastName:  "Jimerson",
		contactInfo: contactInfo{
			email:   "jim@example.com",
			zipCode: "12345", // multiline structs must end with comma, even if their are the last line
		}, // multiline struct comma
	}

	// jimPointer := &jim
	// jimPointer.updateFirstName("jimmy")
	// or (&jim).updateFirstName("jimmy")
	jim.updateFirstName("jimmy") // go shortcut, automatically changes to pointer to jim

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// that doesn't work "as expected" because the person is passed as a value (copy) not a pointer
// func (p person) updateFirstName(firstName string) {
// 	p.firstName = firstName
// }

// jimPointer := &jim
// jimPointer.updateFirstName("jimmy")
func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}
