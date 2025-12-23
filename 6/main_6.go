package main

import "math/rand"

func Random() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Intn(100)
		}
	}()
	return ch
}

func main() {
	k := Random()
	for i := 0; i < 10; i++ {
		println(<-k)
	}
}
