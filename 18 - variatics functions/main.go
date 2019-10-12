package main

import "fmt"

func main() {
	greet := "Hello!"

	sayHi(greet, "Erick", "Alejandro")
}

func sayHi(greet string, names ...string) {
	for _, v := range names {
		fmt.Printf("%s %s. \n", greet, v)
	}
}
