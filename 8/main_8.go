package main

import (
	"fmt"
	"sync"
	"time"
)

type CustomWaitGroup struct {
	sem     chan struct{}
	done    chan struct{}
	mu      sync.Mutex
	counter int
}

func NewCustomWaitGroup(size int) *CustomWaitGroup {
	return &CustomWaitGroup{
		sem:  make(chan struct{}, size),
		done: make(chan struct{}),
	}
}

func (wg *CustomWaitGroup) Add(n int) {
	wg.mu.Lock()
	wg.counter += n
	wg.mu.Unlock()

	for i := 0; i < n; i++ {
		wg.sem <- struct{}{}
	}
}

func (wg *CustomWaitGroup) Done() {
	<-wg.sem
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter--

	if wg.counter < 0 {
		panic("negative WaitGroup counter")
	}

	if wg.counter == 0 {
		close(wg.done)
	}
}

func (wg *CustomWaitGroup) Wait() {
	wg.mu.Lock()
	if wg.counter == 0 {
		wg.mu.Unlock()
		return
	}
	done := wg.done
	wg.mu.Unlock()
	<-done
}

func main() {
	wg := NewCustomWaitGroup(3)
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wg.Done()

			fmt.Println("start", i)
			time.Sleep(time.Second)
			fmt.Println("finish", i)
		}(i)
	}

	wg.Wait()

	fmt.Println("all done")
}
