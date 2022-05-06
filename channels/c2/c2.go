package main

func main() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 13 // 将发送操作放到另外一个 goroutine 里
	}()
	n := <-ch1
	println(n) // 13
}
