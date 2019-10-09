package main

import "fmt"

func main() {
	var names [2]string
	var numbers [2]int

	names[0] = "Erick"
	names[1] = "Alejandro"

	numbers[0] = 1
	numbers[1] = 2

	fmt.Println(names)
	fmt.Println(numbers)
}
