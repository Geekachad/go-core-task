package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	chan1 := FillChan(3)
	chan2 := FillChan(3)
	chan3 := FillChan(3)

	generalChan := MakeGenChan(chan1, chan2, chan3)
	for val := range generalChan {
		fmt.Println(val)
	}

}

func MakeGenChan(channels ...chan int) <-chan int {
	result := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, channel := range channels {
		go func() {
			defer wg.Done()
			for val := range channel {
				result <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func FillChan(n int) chan int {
	ch := make(chan int, n)
	for i := 0; i < n; i++ {
		ch <- rand.Intn(100)
	}
	close(ch)
	return ch
}
