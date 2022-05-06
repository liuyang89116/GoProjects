package main

func main() {
	ch1 := make(chan int)
	ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
	n := <-ch1
	println(n)
}
