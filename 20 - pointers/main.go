package main

import "fmt"

func main() {
	a := 3
	addOne(a)
	fmt.Println(a)
	addOne(a)
	fmt.Println(a)
	reallyAddOne(&a)
	fmt.Println(a)
	reallyAddOne(&a)
	fmt.Println(a)
}

func addOne(v int) {
	v++
	fmt.Println(v)
}

func reallyAddOne(v *int) {
	*v++
	fmt.Println(*v)
}
