package main

import "time"

type signal struct{}

func worker() {
	println("worker is working...")
	time.Sleep(2 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker starts to work...")
		f()
		c <- signal(struct{}{})
	}()

	return c
}

func main() {
	println("init a worker...")
	c := spawn(worker)
	<-c
	println("worker is done!")
}
