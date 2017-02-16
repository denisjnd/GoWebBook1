package main

import "fmt"

func main() {
	// initialize a map
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 4}
	// map has 2 return values. For the second return value, if the key doesn't exist, OK if false. It returns True otherwise
	cSharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the list, its rating is ", cSharpRating)
	} else {
		fmt.Println("We have no rating associated with c#")
	}

	delete(rating, "C") // delete element with key C form the rating map
}
