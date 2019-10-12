package main

import "fmt"

type person struct {
	name string
}

type animal struct {
	name    string
	petType string
}

func (p person) greet() {
	fmt.Printf("Hi! my name is %s. \n", p.name)
}

func (a animal) greet() {
	if a.petType == "dog" {
		fmt.Println("Woff!")
	} else if a.petType == "cat" {
		fmt.Println("Meow")
	} else {
		fmt.Printf("Do nothing...")
	}
}

type entity interface {
	greet()
}

func greet(e entity) {
	fmt.Println("Hi")
	e.greet()
}

func main() {
	friend := person{name: "Alejandro"}
	firulais := animal{name: "Firulais", petType: "dog"}
	mishi := animal{name: "Mishi", petType: "cat"}

	greet(friend)
	greet(firulais)
	greet(mishi)
}
