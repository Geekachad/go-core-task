package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestCustomWaitGroup_Wait(t *testing.T) {
	wg := NewCustomWaitGroup(1)

	wg.Add(1)

	finished := make(chan struct{})

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
		t.Fatal("Wait returned before Done()")
	case <-time.After(100 * time.Millisecond):
		// ожидаем блокировку
	}

	wg.Done()

	select {
	case <-finished:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return after Done()")
	}
}

func TestCustomWaitGroup_MultipleWorkers(t *testing.T) {
	const workers = 5

	wg := NewCustomWaitGroup(workers)

	wg.Add(workers)

	var counter atomic.Int32

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			time.Sleep(10 * time.Millisecond)

			counter.Add(1)
		}()
	}

	wg.Wait()

	if counter.Load() != workers {
		t.Fatalf(
			"expected %d completed workers, got %d",
			workers,
			counter.Load(),
		)
	}
}

func TestCustomWaitGroup_EmptyWait(t *testing.T) {
	wg := NewCustomWaitGroup(1)

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Wait blocked on empty group")
	}
}

func TestCustomWaitGroup_ExtraDoneBlocks(t *testing.T) {
	wg := NewCustomWaitGroup(1)

	wg.Add(1)
	wg.Done()

	finished := make(chan struct{})

	go func() {
		wg.Done()
		close(finished)
	}()

	select {
	case <-finished:
		t.Fatal("extra Done should block")
	case <-time.After(100 * time.Millisecond):
		// ожидаем блокировку
	}
}

func TestCustomWaitGroup_Race(t *testing.T) {
	const workers = 100

	wg := NewCustomWaitGroup(workers)

	wg.Add(workers)

	var result atomic.Int32

	for i := 0; i < workers; i++ {
		go func() {
			result.Add(1)
			wg.Done()
		}()
	}

	wg.Wait()

	if result.Load() != workers {
		t.Fatalf("expected %d, got %d", workers, result.Load())
	}
}
