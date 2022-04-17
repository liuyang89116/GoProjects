package main

import "fmt"

func foo() {
	println("call foo")
	bar()
	println("exit foo")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover panic:", e)
		}
	}()
	println("call bar")
	panic("panic occurs in bar")
	zoo()
	println("exit bar")
}

func zoo() {
	println("call zoo")
	println("exit zoo")
}

func main() {
	println("call main")
	foo()
	println("exit main")

	/*
		call main
		call foo
		call bar
		recover panic: panic occurs in bar
		exit foo
		exit main
	*/
}
