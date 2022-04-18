package main

import "fmt"

type T struct{}

func (t T) M(n int) {
	fmt.Println("M is called. n is", n)
}

func main() {
	var t T
	t.M(1)

	p := &T{}
	p.M(2)
}
