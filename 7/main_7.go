package main

import (
	"fmt"
	"sync"
)

func merge(chans ...chan int) chan int {
	result := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range chans {
		if ch == nil {
			continue
		}
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for v := range c {
				result <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}

func producer(values []int) chan int {
	ch := make(chan int)

	go func() {
		for _, v := range values {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func main() {
	ch1 := producer([]int{1, 2, 3})
	ch2 := producer([]int{10, 11, 12})
	ch3 := producer([]int{100, 101, 102})

	merged := merge(ch1, ch2, ch3)

	for v := range merged {
		fmt.Println(v)
	}
}
