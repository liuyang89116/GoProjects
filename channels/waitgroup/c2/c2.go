package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int { // receive only
	c := make(chan int)

	go func() {
		for {
			fmt.Printf("workder %d received %c\n", id, <-c)
		}
	}()

	return c
}

func chanDemo2() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'B' + i
	}

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo2()
}
