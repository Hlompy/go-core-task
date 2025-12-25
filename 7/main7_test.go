package main

import (
	"reflect"
	"sort"
	"testing"
)

func makeChan(values []int) chan int {
	ch := make(chan int)
	go func() {
		for _, v := range values {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func TestMergeBasic(t *testing.T) {
	ch1 := makeChan([]int{1, 2, 3})
	ch2 := makeChan([]int{10, 11})
	ch3 := makeChan([]int{100})

	out := merge(ch1, ch2, ch3)

	var result []int
	for v := range out {
		result = append(result, v)
	}

	expected := []int{1, 2, 3, 10, 11, 100}
	sort.Ints(result)
	sort.Ints(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
