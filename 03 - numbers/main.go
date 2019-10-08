package main

import "fmt"

func main() {
	var a int
	var b int
	var c int8
	var d int8
	var e int16
	var f int16

	a, b, c, d, e, f = 1, 2, 3, 4, 5, 6

	g := a + b
	h := c + d
	i := e + f

	// casting
	j := a + int(c)
	k := e + int16(d)

	fmt.Println(g, h, i, j, k)

}
