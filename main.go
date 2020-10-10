package main

import "fmt"

// func swap(x, y string) (string, string) {
// 	return y, x
// }

func split(sum int) (a, b int) {
	a = sum * 4 / 9
	b = sum - a
	return
}

func main() {
	// a, b := swap("Hello", "World")
	fmt.Println(split(18))
}
