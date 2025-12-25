package main

import (
	"sync/atomic"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	wg := NewWaitGroup()

	const workers = 5
	wg.Add(workers)

	var completed int32

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&completed, 1)
		}()
	}

	wg.Wait()

	if completed != workers {
		t.Fatalf("expected %d workers completed, got %d", workers, completed)
	}
}

func TestWaitGroupDone(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(3)

	wg.Done()
	wg.Done()
	wg.Done()

	wg.Wait()
}
