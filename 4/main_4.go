package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	slice3 := PresentRightAbsentLeft(slice1, slice2)
	fmt.Println(slice3)
}

func PresentRightAbsentLeft(slice1 []string, slice2 []string) []string {
	result := make([]string, 0, len(slice1))

	set2 := SliceToSet(slice2)

	for _, val1 := range slice1 {
		if _, ok := set2[val1]; !ok {
			result = append(result, val1)
		}
	}
	return result
}

func SliceToSet(slice []string) map[string]struct{} {
	set := map[string]struct{}{}
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}
