package main

import "time"

func main() {
	time.AfterFunc(1*time.Second, func() {
		println("hello world!")
	})

	time.Sleep(2 * time.Second)
}
