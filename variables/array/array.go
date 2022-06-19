package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a)

	for i, v := range a[:] { // or &a
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after for loop, a is", a)
	fmt.Println("after for loop, r is", r)
}
