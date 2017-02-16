package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human     // struct as embedded field
	Skills    // Slice as embedded field
	int       // Built-in type as embedded field
	specialty string
}

func main() {
	// Initialize student Jane
	jane := Student{Human: Human{"Jane", 30, 100}, specialty: "Biology"}
	// access fields
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her specialty is ", jane.specialty)

	// modify value of skill field
	jane.Skills = []string{"Anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new skills")
	// add values to the slice
	jane.Skills = append(jane.Skills, "Physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)

	// modify embedded field of inbuilt type
	jane.int = 3
	fmt.Println("Her prefered number is ", jane.int)
}
