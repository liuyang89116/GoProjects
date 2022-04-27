package main

import "fmt"

func main() {
	var s string = "hello world"
	fmt.Println(s) // hello world

	var sp *string = &s
	fmt.Println(sp) // 0xc000088220
	fmt.Println(&s) // 0xc000088220

	fmt.Println(*sp) // hello world
}
