package main

import "fmt"

func main() {
	var arr [10]int // an array of type int that will contain 10 element
	arr[0] = 20 // array is 0 based
	arr[1] = 39
	fmt.Printf("First value %d \n", arr[0])
	// if i ask for the index that i have not yet assigned any value, it will display the default value which is 0
	fmt.Printf("Last value : %d", arr[9])
}
