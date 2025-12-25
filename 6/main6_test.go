package main

import (
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	r := Random()

	for i := 0; i < 10; i++ {
		select {
		case v := <-r:
			if v < 0 || v >= 100 {
				t.Fatalf("value out of range: %d", v)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatal("timeout: generator did not produce value")
		}
	}
}
