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
			zipCode: 94000,
		},
	}

	//jimPointer := &jim // '&' this operator is getting the variable address and assigning it to a new variable

	// we can gain the same result doing this...
	jim.updateName("jimmy")
	jim.print()
}

// '*' this operator in front a type means that we're expecting the type 'person' from a pointer
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // '*' this operator is getting the value from the pointer address
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
