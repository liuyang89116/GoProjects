package main

import (
	"log"
	"sync"
	"time"
)

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i + 1)
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}

/**
2022/05/07 00:21:22 handle job: 8
2022/05/07 00:21:22 handle job: 2
2022/05/07 00:21:22 handle job: 1
2022/05/07 00:21:24 handle job: 4
2022/05/07 00:21:24 handle job: 3
2022/05/07 00:21:24 handle job: 7
2022/05/07 00:21:26 handle job: 5
2022/05/07 00:21:26 handle job: 6
*/
