package main

import (
	"errors"
	"fmt"
)

var a int = 2020

func checkYear() error {
	err := errors.New("wrong year")

	switch a, err := getYear(); a { // a 和 err 都被遮蔽
	case 2020:
		fmt.Println("it is", a, err)
	case 2021:
		fmt.Println("it is", a)
		err = nil
	}
	fmt.Println("after check, it is", a)

	return err // return 回了 line 11 定义的 err
}

type new int

func getYear() (new, error) {
	var b int16 = 2021
	return new(b), nil
}

func main() {
	err := checkYear() // 返回了 wrong year
	if err != nil {
		fmt.Println("call checkYear error:", err)
		return
	}
	fmt.Println("call checkYear ok")
}
