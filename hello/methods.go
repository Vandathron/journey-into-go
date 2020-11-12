package main

import (
	"errors"
	"fmt"
	"math"
)

// Vertex go
type Vertex struct {
	X, Y float64
}

//Abs returns absolue
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Scale reduces
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Oya does
func Oya() {
	var male ihuman = &gender{"vandathron", 10}

	var female ihuman = &gender{"Sparky", 8}

	male.run()
	male.talk()
	male.changeAge(1, 4)
	male.changeName("vandaton", "van")
	female.run()
	female.talk()
	male.run()
}

type ihuman interface {
	run()
	talk()
	changeAge(currentAge, age int)
	changeName(currentName, name string)
}

type gender struct {
	name string
	age  int
}

func (v gender) run() {
	fmt.Printf("%v with age %v is running \n", v.name, v.age)
}

func (v gender) talk() {
	fmt.Printf("My name is %v \n", v.name)
}

func (v *gender) changeAge(currentAge, age int) {
	if v.age != currentAge {
		err := errors.New("Current age no match")
		println(err.Error())
		return
	}
	v.age = age
}

func (v *gender) changeName(currentName, name string) {
	if v.name != currentName {
		err := errors.New("Current name no match")
		println(err.Error())
		return
	}
	v.name = name
}
