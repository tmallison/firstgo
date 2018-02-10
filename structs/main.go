package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	thomas := person{"Thomas", "Allison"}

	fmt.Println(thomas)
}