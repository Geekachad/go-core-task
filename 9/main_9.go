package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer close(ch1)
		for i := 0; i < 5; i++ {
			ch1 <- uint8(rand.Intn(100))
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()

		for val := range ch2 {
			fmt.Println(val)
		}
	}()

	go func() {
		defer wg.Done()
		CubeFloat(ch1, ch2)
	}()

	wg.Wait()
	fmt.Println("all done")
}

func CubeFloat(ch1 <-chan uint8, ch2 chan<- float64) {
	defer close(ch2)
	for val := range ch1 {
		ch2 <- math.Pow(float64(val), 3)
	}
}
