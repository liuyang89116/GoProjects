package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched() // 让某些协程交出控制权
				// [894 1169 402 486 328 466 480 499 443 512]
			}
		}(i)
	}

	time.Sleep(time.Millisecond) // 否则 main 也是一个协程，里面还没来得及打印，main 就结束了
	fmt.Println(a)
}
