package main

import "fmt"

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2() {
	t.a = 11
}

func main() {
	var t T
	fmt.Println(t.a) // 0

	t.M1()
	fmt.Println(t.a) // 0

	p := &t
	p.M2()
	fmt.Println(t.a) // 11
}
