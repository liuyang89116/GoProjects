package main

import (
	"fmt"
	"time"
)

func worker2(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %d\n", id, <-c)
	}
}

func bufferedChannelDemo() {
	c := make(chan int, 3)

	go worker2(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	time.Sleep(time.Millisecond)
}

func main() {
	bufferedChannelDemo()
}
