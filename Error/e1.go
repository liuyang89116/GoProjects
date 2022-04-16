package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("demo error")
	fmt.Println(err) // demo error

	errWithCtx := fmt.Errorf("index %d is out of bounds", 10)
	fmt.Println(errWithCtx)
}
