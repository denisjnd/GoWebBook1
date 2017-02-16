package main

import "fmt"

type Person struct {
	name string
	age  int
}

// function to compare the age of two people
func Older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {

	var tom Person

	// initialization
	tom.name, tom.age = "Tom", 20

	// initialize two values by format "field:value"
	bob := Person{name: "Bob", age: 30}

	// initialize two values with order
	paul := Person{"Paul", 40}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years \n", tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years \n", tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years \n", bob.name, paul.name, bp_Older.name, bp_diff)
}
