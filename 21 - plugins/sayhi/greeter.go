package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("hello world")
}

//Greeter with custom greet to any name
var Greeter greeting
