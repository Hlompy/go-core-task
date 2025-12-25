package main

import (
	"fmt"
	"sync/atomic"
)

type WaitGroup struct {
	sem   chan struct{}
	count int32
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sem: make(chan struct{}),
	}
}

func (wg *WaitGroup) Add(n int) {
	atomic.AddInt32(&wg.count, int32(n))
}

func (wg *WaitGroup) Done() {
	if atomic.AddInt32(&wg.count, -1) == 0 {
		close(wg.sem)
	}
}

func (wg *WaitGroup) Wait() {
	<-wg.sem
}

func main() {
	wg := NewWaitGroup()

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("worker 1 done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("worker 2 done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("worker 3 done")
	}()

	wg.Wait()
	fmt.Println("all workers done")
}
