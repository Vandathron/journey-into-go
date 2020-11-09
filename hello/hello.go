package main

import (
	"fmt"
	"log"
	"math/cmplx"

	"example.com/greetings"
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
}
