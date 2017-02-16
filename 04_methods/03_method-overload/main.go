package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // anonymous field
	school string
}

type Employee struct {
	Human
	company string
}

// define a method in Human
func (h *Human) SayHi() {
	fmt.Printf("Hi, i am %s, you can call on %s \n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, am %s, i work at %s. Call me on %s \n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"Mark", 20, "234-8878-653"}, "MIT"}
	sam := Employee{Human{"Sam", 39, "41-576-8874-453"}, "Google"}

	mark.SayHi()
	sam.SayHi()
}
