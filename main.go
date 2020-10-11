package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y int
}

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

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

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

	// fmt.Print("Go runs on ...")
	// switch os := runtime.GOOS; os {
	// case "Linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	fmt.Printf("%s.\n", os)
	// }

	// fmt.Println("counting")
	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println("done")

	// i, j := 42, 2701
	// p := &i
	// fmt.Println(*p)
	// *p = 21
	// fmt.Println(i)

	// p = &j
	// *p = *p / 37
	// fmt.Println(j)

	// v := Vertex{1, 2}
	// v.X = 4
	// fmt.Println(v)
	// fmt.Println(v.X)

	// v := Vertex{1, 2}
	// p := &v
	// (*p).X = 1e9
	// fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
