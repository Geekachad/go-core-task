package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := 10
	originalSlice := make([]int, num)
	for i := 0; i < num; i++ {
		originalSlice[i] = rand.Int()
	}
	fmt.Println(originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println(evenSlice)

	randNum := rand.Int()
	newSlice := addElements(originalSlice, randNum)
	fmt.Println(newSlice)

	sliceCopy := copySlice(originalSlice)
	fmt.Println(sliceCopy)

	sliceWithRemoveElement := removeElement(originalSlice, 2)
	fmt.Println(sliceWithRemoveElement)
}

func sliceExample(slice []int) []int {
	var result []int
	for _, value := range slice {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

func addElements(slice []int, num int) []int {
	return append(slice, num)
}

func copySlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	return result
}

func removeElement(slice []int, idx int) []int {
	if idx < 0 || idx >= len(slice) {
		return nil
	}
	result := make([]int, idx, len(slice)-1)
	copy(result, slice[:idx])
	result = append(result, slice[idx+1:]...)
	return result
}
