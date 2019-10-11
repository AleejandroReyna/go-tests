package main

import "fmt"

func main() {
	//increment
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	//decrement
	for i := 3; i >= 0; i-- {
		fmt.Println(i)
	}

	//break
	for i := 0; i >= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	//practice
	matriz := [3][3]int{}
	initial := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			initial++
			matriz[i][j] = initial
		}
	}

	fmt.Println(matriz)
}
