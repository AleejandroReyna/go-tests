package main

import "fmt"

type greeting string

func (g greeting) Greet(greet string, name string) {
	fmt.Printf("%s %s", greet, name)
}

//Greeter with custom greet to any name
var Greeter greeting
