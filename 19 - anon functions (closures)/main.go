package main

import "fmt"

func main() {
	func() {
		fmt.Println("Auto execute")
	}()

	anon := func() {
		fmt.Println("Anon func")
	}
	anon()
	firstSecuence := secuence()
	secondSecuence := secuence()

	fmt.Println(firstSecuence())
	fmt.Println(firstSecuence())
	fmt.Println(firstSecuence())

	fmt.Println(secondSecuence())
	fmt.Println(secondSecuence())
	fmt.Println(secondSecuence())
}

func secuence() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
