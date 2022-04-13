package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	c1, c2 := generate(), generate()
	for {
		select {
		case n := <-c1:
			fmt.Println("Receive from c1: ", n)
		case n := <-c2:
			fmt.Println("Receive from c2: ", n)
		}
	}
}
