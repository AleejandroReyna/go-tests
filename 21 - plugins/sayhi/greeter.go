package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("hello world")
}

//Greeter plugin generate  use the flag -buildmode=plugin and -o path/filename.so -o path/plugin.go
var Greeter greeting
