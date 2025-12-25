package main

import (
	"reflect"
	"testing"
)

func TestCubePipeline(t *testing.T) {
	in := make(chan uint8)

	out := cubePipeline(in)

	go func() {
		values := []uint8{1, 2, 3, 4, 5}
		for _, v := range values {
			in <- v
		}
		close(in)
	}()

	var result []float64
	for v := range out {
		result = append(result, v)
	}

	expected := []float64{
		1,
		8,
		27,
		64,
		125,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
