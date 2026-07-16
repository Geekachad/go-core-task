package main

import (
	"testing"
	"time"
)

func TestMakeRandChan_GeneratesValues(t *testing.T) {
	randChan := make(chan int)
	done := make(chan struct{})
	n := 10

	go MakeRandChan(randChan, n, done)

	for i := 0; i < 5; i++ {
		select {
		case val, ok := <-randChan:
			if !ok {
				t.Fatal("The channel closed unexpectedly")
			}
			if val < 0 || val >= n {
				t.Errorf("Incorrect value received: %d, expected from 0 to %d", val, n-1)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatal("Timeout: The function did not send a value to the channel")
		}
	}

	close(done)
}

func TestMakeRandChan_ClosesOnDone(t *testing.T) {
	randChan := make(chan int)
	done := make(chan struct{})
	n := 100

	go MakeRandChan(randChan, n, done)

	close(done)

	timeout := time.After(200 * time.Millisecond)
	for {
		select {
		case _, ok := <-randChan:
			if !ok {
				// Канал успешно закрылся, тест пройден
				return
			}
		case <-timeout:
			t.Fatal("Timeout: Channel randChan was not closed after closing done")
		}
	}
}
