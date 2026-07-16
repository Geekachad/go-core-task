package main

import (
	"testing"
	"time"
)

func TestFillChan_CorrectLength(t *testing.T) {
	n := 5
	ch := FillChan(n)

	count := 0
	for val := range ch {
		count++
		if val < 0 || val >= 100 {
			t.Errorf("Value out of range [0, 100): %d", val)
		}
	}

	if count != n {
		t.Errorf("expected: %d, got: %d", n, count)
	}
}

func TestMakeGenChan_Success(t *testing.T) {
	ch1 := FillChan(3)
	ch2 := FillChan(2)

	resChan := MakeGenChan(ch1, ch2)

	expectedCount := 5
	actualCount := 0

	timeout := time.After(500 * time.Millisecond)

	for {
		select {
		case _, ok := <-resChan:
			if !ok {
				goto CheckCount
			}
			actualCount++
		case <-timeout:
			t.Fatal("Timeout: Timeout waiting for data or channel closing")
		}
	}

CheckCount:
	if actualCount != expectedCount {
		t.Errorf("expected: %d, got: %d", expectedCount, actualCount)
	}
}

func TestMakeGenChan_NoInputs(t *testing.T) {
	resChan := MakeGenChan()

	select {
	case _, ok := <-resChan:
		if ok {
			t.Error("The channel should be empty and close immediately")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Timeout: The function hung due to the lack of input channels")
	}
}
