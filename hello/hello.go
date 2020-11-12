package main

import (
	"fmt"
	"log"
	"math/cmplx"
	"runtime"

	"example.com/greetings"
	"example.com/types"
)

var i, j int = 1, 2
var name string = "tosin"
var (
	ToBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Tosin", "Jahbless", "Dele"}

	//Request a greeting message.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	sayhi(calc(10, 5))
	// whileLoop()
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	types.Mapping()

	v := &Vertex{3, 4}
	x := Vertex{2, 2}
	x.Scale(10)
	v.Scale(10)
	fmt.Println(v.Abs(), x.Abs())

	Oya()
}

func sayhi(x, y, z int) {
	fmt.Println(addNum(23, 10))
	fmt.Println(x, y, z)
}

func addNum(x, y int) int {
	return x + y
}

func calc(x, y int) (int, int, int) {
	add := x + y
	min := x - y
	mul := x * y
	return add, min, mul
}

//while loop

func whileLoop() {
	sum := 0
	for sum < 100 {
		println(sum)
		sum += 2
	}
	ifFunction(3)

	pow := []int{2, 5, 3, 5, 6, 5}

	for _, v := range pow {
		println(v)
	}
}

func ifFunction(v int) {
	if v == 2 {
		println("Hell0")
	}

	switchCase()
}

func switchCase() {
	switch v := runtime.GOOS; v {
	case "darwin":
		fmt.Println("Darwing")
		break
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd
		fmt.Println("Windows")
	}
}

func pointers() {
	var p *int
	println(p)

	v := 20

	var pointerOfV = &v
	println(*pointerOfV * 10)

	println(p)

	a := 10
	var pointerOfA = &a
	b := 30
	var pointerOfB = &b
	pointerOfA = pointerOfB
	println(a)
	var g = &a
	*g = 20
	println(a)
	println(pointerOfA)
}
