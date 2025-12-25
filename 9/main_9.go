package main

import "fmt"

func cubePipeline(in <-chan uint8) <-chan float64 {
	out := make(chan float64)

	go func() {
		defer close(out)
		for v := range in {
			f := float64(v)
			out <- f * f * f
		}
	}()

	return out
}

func main() {
	in := make(chan uint8)

	out := cubePipeline(in)

	go func() {
		for i := uint8(1); i <= 5; i++ {
			in <- i
		}
		close(in)
	}()

	for v := range out {
		fmt.Printf("Value %v of Type %T\n", v, v)
	}
}
