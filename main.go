package main

import (
	"fmt"
	"math"
	"runtime"
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

// var i int = 5

func sqrt(i float64) string {
	if i < 0 {
		return sqrt(-i) + "i"
	}
	return fmt.Sprint(math.Sqrt(i))
}

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}

func sqrt22(x float64) float64 {
	z := 1.0
	for math.Abs(z*z-x) > 0.0001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	// a, b := swap("Hello", "World")
	// var i int
	// fmt.Println(i, c, Python, Java)

	// x := 3
	// y := 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// z := uint(f)
	// fmt.Println(x, y, f, z)
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// 	fmt.Println(sum)
	// }
	// fmt.Println(sqrt(2), sqrt(-4))
	// fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	// fmt.Println(sqrt22(3))
	fmt.Print("Go runs on ...")
	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
