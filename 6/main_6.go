package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randChan := make(chan int)
	done := make(chan struct{})
	go MakeRandChan(randChan, 100, done)

	for i := 0; i < 10; i++ {
		fmt.Println(<-randChan)
	}
	close(done)
}

// MakeRandChan генерирует случайное число.
//
// randChan канал с сгенерированным числом.
// n максимальное сгенерированное число.
// done канал для завершения.
func MakeRandChan(randChan chan<- int, n int, done <-chan struct{}) {

	for {
		select {
		case <-done:
			close(randChan)
			return
		case randChan <- rand.Intn(n):
		}
	}
}
