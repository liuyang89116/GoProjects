package main

func main() {
	//ch := make(chan int, 1)
	//n := <-ch // 缓冲区没东西，所以会被挂起。
	//// fatal error: all goroutines are asleep - deadlock!
	//println(n)

	ch2 := make(chan int, 1)
	ch2 <- 17 // 没问题，向 ch 发送 17
	ch2 <- 27 // 缓冲区满了，发东西也会被挂起
}
