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

	var ok1, ok2 bool
	for {
		select {
		case x := <-ch1:
			ok1 = true
			fmt.Println(x)
		case x := <-ch2:
			ok2 = true
			fmt.Println(x)
		}

		if ok1 && ok2 {
			break
		}
	}
	fmt.Println("program end")
}

/**
5
0
0
0
0
0
7
*/
