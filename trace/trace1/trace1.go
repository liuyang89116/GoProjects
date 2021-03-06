package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var goroutineSpace = []byte("goroutine ")
var mu sync.Mutex
var m = make(map[uint64]int)

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("cannot find caller!")
	}

	fn := runtime.FuncForPC(pc) // 拿到 func
	name := fn.Name()
	gid := curGoroutineId()

	mu.Lock()
	indents := m[gid]    // 获取当前 gid 对应的缩进层次
	m[gid] = indents + 1 // 缩进层次 + 1 后存入 map
	mu.Unlock()
	printTrace(gid, name, "->", indents+1)

	return func() {
		mu.Lock()
		indents := m[gid]    // 获取当前 gid 对应的缩进层次
		m[gid] = indents - 1 // 缩进层次 -1 后存入 map
		mu.Unlock()
		printTrace(gid, name, "<-", indents)
	}
}

func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "   "
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}

func curGoroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	//  parse the 4706 out of "goroutine 4706 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine Id out of %q: %v", b, err))
	}

	return n
}

func A1() {
	defer Trace()()
	B1()
}

func B1() {
	defer Trace()()
	C1()
}

func C1() {
	defer Trace()()
	D()
}

func D() {
	defer Trace()()
}

func A2() {
	defer Trace()()
	B2()
}

func B2() {
	defer Trace()()
	C2()
}

func C2() {
	defer Trace()()
	D()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()

	A1()
	wg.Wait()
}

/*
	g[00001]:   ->main.A1
	g[00001]:      ->main.B1
	g[00001]:         ->main.C1
	g[00006]:   ->main.A2
	g[00006]:      ->main.B2
	g[00006]:         ->main.C2
	g[00001]:            ->main.D
	g[00001]:            <-main.D
	g[00001]:         <-main.C1
	g[00001]:      <-main.B1
	g[00001]:   <-main.A1
	g[00006]:            ->main.D
	g[00006]:            <-main.D
	g[00006]:         <-main.C2
	g[00006]:      <-main.B2
	g[00006]:   <-main.A2
*/
