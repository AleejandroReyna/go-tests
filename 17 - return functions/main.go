package main

import "fmt"

func main() {
	base := 4
	height := 2
	numbers := []int{2, 4, 6, 10, 12}
	area := area(base, height)
	fmt.Println(area)
	min, max := minAndMax(numbers)
	fmt.Println(min, max)
}

func area(base, height int) int {
	return base * height
}

func minAndMax(numbers []int) (min, max int) {
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}

	min = max

	for _, v := range numbers {
		if v < min {
			min = v
		}
	}

	return min, max
}
