package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// same as contactInfo contactInfo
	contactInfo
}

func main() {
	// different ways to create a person

	// relies on order of fields for assignment
	// snack := person{"Snack", "Mullen"}

	// named property/value pairs
	// doyle := person{firstName: "Osc", lastName: "Doyle"}
	// fmt.Println(doyle)

	// zero-value assigned to fields (empty strings in this case)
	// var snack person
	// snack.firstName = "Snack"
	// snack.lastName = "Mullen"
	// fmt.Printf("%+v", snack)

	snack := person{
		firstName: "Snack",
		lastName:  "Mullen",
		contactInfo: contactInfo{
			email:   "snack@test.com",
			zipCode: 12345,
		},
	}

	// Get memory address of "snack" variable
	// snackPointer := &snack
	// snackPointer.updateFirstName("Luck")

	snack.updateFirstName("Luck")
	snack.print()
}

// *person is a pointer that points at a person
func (p *person) updateFirstName(name string) {
	// give value that is at the pointerToPerson memory address
	(*p).firstName = name
}

// Does not work because pass by value
// func (p person) updateFirstName(name string) {
// 	p.firstName = name
// }

// function with receiver of type person
func (p person) print() {
	fmt.Printf("%+v", p)
}
