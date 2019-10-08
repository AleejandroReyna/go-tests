package main

import "fmt"

func main() {
	var name, secondName string
	name = "Erick"
	secondName = "Alejandro"
	lastName, anotherLastName := "Alvarez", "Reyna"

	fmt.Println(name, secondName, lastName, anotherLastName)
}
