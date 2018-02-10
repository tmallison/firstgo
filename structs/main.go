package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	thomas := person{"Thomas", "Allison"}

	thomasPointer := &thomas
	thomasPointer.updateName("Tommy")
	fmt.Println(thomas)

	thomas.updateName("Tim")
	fmt.Println(thomas)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}