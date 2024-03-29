package main

import (
	"fmt"
	"os"
	"plugin"
)

//Greeter interface
type Greeter interface {
	Greet(name string)
}

func main() {

	plug, err := plugin.Open("./sayminame/sayminame.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	greeter.Greet("Alejandro")
}
