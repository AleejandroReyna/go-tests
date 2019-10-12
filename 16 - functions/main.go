package main

import "fmt"

func main() {
	name := "Alejandro"
	sayHi(name)
}

func sayHi(name string) {
	fmt.Printf("Hi %s.\n", name)
}
