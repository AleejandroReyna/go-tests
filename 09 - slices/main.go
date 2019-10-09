package main

import "fmt"

func main() {
	var names []string
	otherNames := []string{
		"Rodolfo",
		"Edgar",
	}
	anotherNames := make([]string, 0)

	//names[1] = "Alejandro" <-- Incorrect
	names = append(names, "Erick")
	names = append(names, "Alvarez")

	fmt.Println(names)
	fmt.Printf("Size: %d. \n Capacity: %d. \n", len(names), cap(names))

	anotherNames = append(anotherNames, "Sucely")
	anotherNames = append(anotherNames, "Pamela")

	fmt.Println(anotherNames)
	fmt.Println(otherNames)
}
