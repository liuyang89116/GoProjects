package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		time.Sleep(7 * time.Second)
		ch2 <- 7
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println(x)
			}
		case x, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println(x)
			}
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("program end")
}

/**
5
7
program end
*/
