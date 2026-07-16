package main

import (
	"math"
	"testing"
)

func TestCubeFloat(t *testing.T) {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go func() {
		defer close(ch1)

		values := []uint8{2, 3, 5}

		for _, v := range values {
			ch1 <- v
		}
	}()

	go CubeFloat(ch1, ch2)

	expected := []float64{
		8,
		27,
		125,
	}

	var result []float64

	for val := range ch2 {
		result = append(result, val)
	}

	if len(result) != len(expected) {
		t.Fatalf("expected %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Fatalf("expected %v, got %v", expected[i], result[i])
		}
	}
}

func TestCubeFloat_ChannelClosed(t *testing.T) {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go func() {
		defer close(ch1)

		ch1 <- 4
	}()

	go CubeFloat(ch1, ch2)

	for range ch2 {
	}

	_, ok := <-ch2

	if ok {
		t.Fatal("expected ch2 to be closed")
	}
}

func TestCubeFloat_EmptyInput(t *testing.T) {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	close(ch1)

	go CubeFloat(ch1, ch2)

	for range ch2 {
		t.Fatal("expected no values")
	}
}

func TestCubeFloat_LargeValue(t *testing.T) {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go func() {
		defer close(ch1)
		ch1 <- 10
	}()

	go CubeFloat(ch1, ch2)

	result := <-ch2

	expected := math.Pow(10, 3)

	if result != expected {
		t.Fatalf("expected %.0f, got %.0f", expected, result)
	}
}

func TestPipeline(t *testing.T) {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	input := []uint8{1, 2, 3, 4}

	go func() {
		defer close(ch1)

		for _, v := range input {
			ch1 <- v
		}
	}()

	go CubeFloat(ch1, ch2)

	var result []float64

	for val := range ch2 {
		result = append(result, val)
	}

	expected := []float64{
		1,
		8,
		27,
		64,
	}

	if !equalFloatSlices(result, expected) {
		t.Fatalf(
			"expected %v, got %v",
			expected,
			result,
		)
	}
}

func equalFloatSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
