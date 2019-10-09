package main

import "fmt"

func main() {
	day := 2

	switch day {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fmt.Println("< 4")
	case 4:
		fmt.Println("= 4")
	default:
		fmt.Println("> 4")
	}
}
