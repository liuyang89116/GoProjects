package main

import (
	"fmt"
	"time"
)

func worker3(id int, c chan int) {
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d received %d\n", id, n)
	//}
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func channelClosedDemo() {
	c := make(chan int, 3)

	go worker3(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	close(c)

	time.Sleep(time.Millisecond)
}

func main() {
	channelClosedDemo()
}
