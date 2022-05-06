package main

import (
	"sync"
	"time"
)

func consumer(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}

/**
最基本的生产者消费者模型
*/
func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consumer(ch)
		wg.Done()
	}()

	wg.Wait()
}
