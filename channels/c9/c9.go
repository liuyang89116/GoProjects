package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(c chan<- int) {
	var i int = 1
	for {
		time.Sleep(2 * time.Second)
		ok := trySend(c, i)

		if !ok {
			fmt.Printf("[producer]: try send [%d] to channel, but it's full!\n", i)
			continue
		}

		fmt.Printf("[producer]: send [%d] to channel.\n", i)
		i++
	}
}

func consumer(c <-chan int) {
	for {
		i, ok := tryConsume(c)
		if !ok {
			fmt.Println("[consumer]: try to consume from channel, but it's empty!")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("[consumer]: consume %d from channel.\n", i)
		if i >= 3 {
			fmt.Println("[consumer]: exit")
			return
		}
	}
}

func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

func tryConsume(c <-chan int) (int, bool) {
	select {
	case i := <-c:
		return i, true
	default:
		return 0, false
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 3)
	wg.Add(2)
	go func() {
		producer(c)
		wg.Done()
	}()

	go func() {
		consumer(c)
		wg.Done()
	}()

	wg.Wait()
}

/**
[consumer]: try to consume from channel, but it's empty!
[consumer]: try to consume from channel, but it's empty!
[consumer]: try to consume from channel, but it's empty!
[producer]: send [1] to channel.
[consumer]: consume 1 from channel.
[consumer]: try to consume from channel, but it's empty!
[consumer]: try to consume from channel, but it's empty!
[producer]: send [2] to channel.
[consumer]: consume 2 from channel.
[consumer]: try to consume from channel, but it's empty!
[consumer]: try to consume from channel, but it's empty!
[producer]: send [3] to channel.
[consumer]: consume 3 from channel.
[consumer]: exit
[producer]: send [4] to channel.
[producer]: send [5] to channel.
[producer]: send [6] to channel.
[producer]: try send [7] to channel, but it's full!
[producer]: try send [7] to channel, but it's full!
[producer]: try send [7] to channel, but it's full!
[producer]: try send [7] to channel, but it's full!
[producer]: try send [7] to channel, but it's full!
[producer]: try send [7] to channel, but it's full!
[producer]: try send [7] to channel, but it's full!
[producer]: try send [7] to channel, but it's full!
*/
