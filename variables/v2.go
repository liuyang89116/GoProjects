package main

import "fmt"

func main() {
	var arr1 = [...]int{
		5: 20, // set 6th element to 20
	}
	fmt.Println(arr1) // [0 0 0 0 0 20]

	var nums = []int{1, 2, 3, 4, 5, 6}
	nums = append(nums, 7)
	fmt.Println(nums)

	arr2 := [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}
	sl := arr2[3:7:9] // [44 55 66 77]
	fmt.Println(sl)
	sl[0] = 0
	fmt.Println(arr2) // [11 22 33 0 55 66 77 88 99 100]

	m := make(map[int]string)
	m[1] = "Bob"
	m[6] = "Amy"
	v, ok := m[6]
	if !ok {
		fmt.Println("doesn't have key 6")
		return
	}
	fmt.Println("the value of key 6 is", v)
}
