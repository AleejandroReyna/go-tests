package main

import "fmt"

func main() {
	a, b, c, d := 5, 10, 15, 20
	years := 21

	if a > b {
		fmt.Println("${a} > b")
	}

	if a < b && c < d {
		fmt.Println("a < b & c < d")
	}

	if !(a > b) {
		fmt.Println("a < b")
	}

	if true {
		fmt.Println("true")
	}

	if 5 != 10 {
		fmt.Println("not equals")
	} else {
		fmt.Println("Equals")
	}

	if years < 14 {
		fmt.Println("kid")
	} else if years < 19 {
		fmt.Println("Young")
	} else if years < 45 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Old")
	}

	fmt.Println("End")
}
