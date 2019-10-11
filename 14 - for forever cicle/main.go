package main

import "fmt"

func main() {
	limit := 0
	for {
		limit++
		fmt.Println("Hi!")

		if limit == 10 {
			break
		}
	}
}
