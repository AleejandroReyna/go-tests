package main

import "fmt"

func main() {
	numbers := []int8{0, 1, 2, 3, 4, 5}
	names := map[int]string{0: "Erick", 1: "Alejandro"}
	quote := "hello world"
	for i, v := range numbers {
		fmt.Println(i, v)
	}

	for _, v := range names {
		fmt.Println(v)
	}

	for i, v := range quote {
		fmt.Println(i, string(v))
	}
}
