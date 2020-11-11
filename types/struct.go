package types

import (
	"fmt"
)

type vertex struct {
	X int
	Y int
}

// PrintStruct returns void
func PrintStruct() {
	fmt.Println("ye")
	fmt.Println(vertex{1, 2})
	v := vertex{4, 3}
	var pointer = &v.X
	*pointer = 30
	fmt.Println(v.X)
	fmt.Println(v)
	fmt.Println(pointer)
}

// Mapping deos
func Mapping() {
	v := make(map[int]vertex)
	v[1] = vertex{X: 3}
	defer fmt.Println(v[1])
	delete(v, 1)
	println("herwrs")
	add := func(x, y int) int {
		return x + y
	}

	sub := func(x, y int) int {
		return x - y
	}

	sub = add

	mul := func(x, y int) int {
		return x * y
	}
	runCal(sub)
	runCal(add)
	runCal(mul)
}

func add(x, y int) int {
	return x + y
}
func runCal(fn func(int, int) int) {
	println(fn(2, 9))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
