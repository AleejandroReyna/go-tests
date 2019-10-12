package main

import "fmt"

type greeter string

func (g greeter) Greet(name string) {
	fmt.Printf("Hi! %s. \n", name)
}

//Greeter for generate
var Greeter greeter
