package main

import "fmt"

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}

	// 3 2 1 0
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}

	// 3 2 1 0
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	// 4 4 4 4
}

func main() {
	fmt.Println("foo1 result:")
	foo1()
	fmt.Println("\nfoo2 result:")
	foo2()
	fmt.Println("\nfoo3 result:")
	foo3()
}
