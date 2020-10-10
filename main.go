package main

import (
	"fmt"
	"math"
)

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func split(sum int) (a, b int) {
// 	a = sum * 4 / 9
// 	b = sum - a
// 	return
// }

// var c, Python, Java bool

func main() {
	// a, b := swap("Hello", "World")
	// var i int
	// fmt.Println(i, c, Python, Java)
	x := 3
	y := 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, f, z)
}
