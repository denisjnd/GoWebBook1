package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("emit some error if any")
	if err != nil {
		fmt.Print(err)
	}
}
