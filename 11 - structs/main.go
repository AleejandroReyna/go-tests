package main

import "fmt"

type Person struct {
	Name      string
	Years     uint8
	LastNames []string
}

func main() {
	var person Person
	person.Name = "Alejandro"
	person.Years = 26
	person.LastNames = []string{"Alvarez", "Reyna"}

	secondPerson := Person{}
	secondPerson.Name = "Erick"
	secondPerson.Years = 25

	thirdPerson := Person{
		Name:  "Jabi",
		Years: 22,
	}

	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(secondPerson)
	fmt.Println(thirdPerson)
}
