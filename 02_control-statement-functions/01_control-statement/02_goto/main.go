package main

import "fmt"

func main() {
	// infinite loop using goto keyword
	i := 0
	Here: // label is case sensitive and it ends with a :
	fmt.Println(i)
	i++
	goto Here
}
