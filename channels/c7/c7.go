package main

import (
	"fmt"
	"sync"
)

type counter struct {
	ch  chan int
	num int
}

func NewCounter() *counter {
	cnt := &counter{
		ch: make(chan int),
	}
	go func() {
		for {
			cnt.num++
			cnt.ch <- cnt.num
		}
	}()
	return cnt
}

func (cnt *counter) Increase() int {
	return <-cnt.ch
}

func main() {
	cnt := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cnt.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait() // 不加这个直接 main goroutine 直接进行完退出了
}
