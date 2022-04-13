package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var c1, c2 = generate(), generate()     // 两个 channel 来生成 num
	var worker chan<- int = createWorker(0) // 一个 worker 来 consume num，把 c1, c2 产生的 num 读进来

	stopTime := time.After(10 * time.Second) // return as a chan
	tick := time.Tick(time.Second)

	// create 一个 slice 用来存放生成的 num，以防生成的太快，consume 不过来
	var arr []int
	for {
		var activeWorker chan<- int
		var activeValue int

		// 如果 slice 里面是有 num 的，那就可以 consume
		if len(arr) > 0 {
			activeWorker = worker
			activeValue = arr[0]
		}

		// 不同的方式可以用来调度
		select {
		case n := <-c1:
			arr = append(arr, n)
		case n := <-c2:
			arr = append(arr, n)
		case activeWorker <- activeValue:
			fmt.Printf("value %d is removing from the buffer queue\n", activeValue)
			arr = arr[1:] // consume 了 activeValue，arr 取剩下的值
		case <-tick:
			fmt.Println("current arr size:", len(arr))
		case <-stopTime:
			fmt.Println("Force exit")
			return
		}
	}
}

// 创建 worker，用来 consumer chan 里的 num
func createWorker(id int) chan<- int {
	w := make(chan int)
	go worker(id, w)
	return w
}

func worker(id int, w chan int) {
	for n := range w {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

// 生成 num 放到 chan 里
func generate() <-chan int {
	out := make(chan int)
	num := 0
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- num
			num++
		}
	}()
	return out
}
