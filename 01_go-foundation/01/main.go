package main

import "fmt"

/*
To change one character in a string, you must convert it to a byte first
*/
func main() {
	var s, w = "Hello", " World"
	c := []byte(s) // convert string to byte type
	c[0] = 'K'
	s2 := string(c) // convert back to string type
	fmt.Printf("%s \n", s2)
	s3 := s2 + w
	fmt.Printf("%s \n", s3)
}
