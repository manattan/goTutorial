package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/wc"
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

var fl = []int{1, 2, 4, 8, 16, 32, 64, 128}
var m map[string]Vertex

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	arr := strings.Fields(s)
	for i := 0; i < len(arr); i++ {
		_, ok := res[arr[i]]
		if ok {
			res[arr[i]] += 1
		} else {
			res[arr[i]] = 1
		}
	}
	return res
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

	// fmt.Println(v1, p, v2, v3)

	// var a [2]string
	// a[0] = "hello"
	// a[1] = "world"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	// primes := [6]int{1, 2, 3, 5, 5, 2}
	// fmt.Println(primes)

	// primes := [6]int{2, 3, 5, 8, 13, 21}
	// var s []int = primes[:4]
	// fmt.Println(s)
	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)

	// // Slice the slice to give it zero length.
	// s = s[:0]
	// printSlice(s)

	// // Extend its length.
	// s = s[:4]
	// printSlice(s)

	// // Drop its first two values.
	// s = s[2:]
	// printSlice(s)

	// var s []int
	// fmt.Println(s, len(s), cap(s))
	// if s == nil {
	// 	fmt.Println("nil!")
	// }

	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }

	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "O"

	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", board[i])
	// }

	// var s []int
	// printSlice(s)
	// s = append(s, 1)
	// printSlice(s)
	// s = append(s, 2, 3, 4)
	// printSlice(s)

	// for i, j := range fl {
	// 	fmt.Printf("2の%d乗は%dです\n", i, j)
	// }

	// ali := make([]int, 10)
	// for i := range ali {
	// 	ali[i] = 1 << uint(i)
	// }
	// for _, value := range ali {
	// 	fmt.Println(value)
	// }

	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{40, 20}
	// fmt.Println(m["Bell Labs"])

	// jk := make(map[string]int)
	// jk["Answer"] = 42
	// fmt.Println("The value: ", jk["Answer"])
	// jk["Answer"] = 48
	// fmt.Println("The value: ", jk["Answer"])
	// v, ok := jk["Answer"]
	// fmt.Println("The value: ", v, "Present?", ok)

	wc.Test(WordCount)
}
