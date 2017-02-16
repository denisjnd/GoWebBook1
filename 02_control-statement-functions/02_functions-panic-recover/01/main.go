package main

import "os"

func init() {
	var user = os.Getenv("USER")
	if user == "" {
		panic("no value for $USER")
	}
}
