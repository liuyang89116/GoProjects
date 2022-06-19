package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Println("Hello from goroutine ", i)
			}
		}(i)
	}

	time.Sleep(time.Second) // 否则 main 也是一个协程，里面还没来得及打印，main 就结束了
}
