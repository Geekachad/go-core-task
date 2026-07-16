package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, result := Intersection(a, b)
	fmt.Println(ok, result)
}

func Intersection(a []int, b []int) (bool, []int) {
	result := make([]int, 0, len(a))
	setB := SliceToSet(b)
	for _, val := range a {
		if _, ok := setB[val]; ok {
			result = append(result, val)
		}
	}
	if len(result) == 0 {
		return false, result
	}
	return true, result
}

func SliceToSet(slice []int) map[int]struct{} {
	set := make(map[int]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}
