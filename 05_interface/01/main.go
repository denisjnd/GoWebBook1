package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, i am %s, you can call me on %s \n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("la la la la ...", lyrics)
}

// method overload
func (e Employee) SayHi() {
	fmt.Printf("Hi, i am %s, i work at %s. You can call me on %s \n", e.name, e.company, e.phone)
}

// interface Men implemented by Human, Student, Employee
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 20, "2234355432"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 23, "4456675345"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 34, "76889065633"}, "Google", 10000}
	tom := Employee{Human{"Tom", 35, "22322323445"}, "Apple", 15000}

	// define interface i
	var i Men

	// i can store student
	i = mike
	fmt.Println("This is Mike, a Student: ")
	i.SayHi()
	i.Sing("Feb Rain")

	// i can store employee
	i = tom
	fmt.Println("This is Tom, an Employee: ")
	i.SayHi()
	i.Sing("Born to be wild")

	// slice of Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	// these 3 elements are different types but the will all implement interface Men
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}
