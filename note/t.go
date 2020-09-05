package main

import (
	"fmt"
	"sync"
	"time"
)

type P struct {
	P1 string
	P2 int
}

func (p *P) Reset() {
	p.P1 = ""
	p.P2 = 0
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go func(index int) {
			cond.L.Lock()
			cond.Wait()
			fmt.Println(index, "awaken")
			cond.L.Unlock()
		}(i)
	}

	cond.Signal()
	time.Sleep(time.Second * 1)
	cond.Signal()
	time.Sleep(time.Second * 1)
	cond.Broadcast()
	for {
		time.Sleep(time.Second * 1)
	}

	//runtime.GOMAXPROCS(1)
	//pool := sync.Pool{}
	//pool.New = func() interface{} {
	//	return &P{}
	//}
	//// head 1000
	//// val  1000
	//// tail 0001
	//for i := 0; i < 8; i++ {
	//	pool.Put(&P{P1: "1", P2: i})
	//}
	//for i := 0; i < 8; i++ {
	//	pool.Get()
	//}
	//for i := 0; i < 16; i++ {
	//	pool.Put(&P{P1: "1", P2: i})
	//}
	//for i := 0; i < 16; i++ {
	//	pool.Get()
	//}
	//pool.Get()
	//for i := 0; i < 10; i++ {
	//	//go func(index int) {
	//	//p1 := pool.Get().(*P)
	//	//p1.P1 = "hello"
	//	//p1.P2 = i
	//	//
	//	//p1.Reset()
	//	pool.Put(&P{P1: "1", P2: i})
	//	//}(i)
	//}
	//for i := 0; i < 10; i++ {
	//	go func(index int) {
	//		size := runtime.GOMAXPROCS(0)
	//		print(fmt.Sprintf("%d grouotine: %d\n", index, size))
	//	}(i)
	//}
	select {}
}
