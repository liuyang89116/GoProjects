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

	t.M2()           // 实际上内部做了转换： 变成了 (&t).M2()
	fmt.Println(t.a) // 11
}
