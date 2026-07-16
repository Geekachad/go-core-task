package main

import (
	"maps"
	"slices"
	"testing"
)

func TestSliceToSet(t *testing.T) {
	slice := []int{65, 3, 58}
	want := map[int]struct{}{
		65: {},
		3:  {},
		58: {},
	}

	result := SliceToSet(slice)

	if !maps.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestSliceToSet_EmptySlice(t *testing.T) {
	slice := []int{}
	want := map[int]struct{}{}

	result := SliceToSet(slice)

	if !maps.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestSliceToSet_NilSlice(t *testing.T) {
	var slice []int
	want := map[int]struct{}{}

	result := SliceToSet(slice)

	if !maps.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	want := []int{3, 64}

	ok, result := Intersection(a, b)

	if !slices.Equal(result, want) || !ok {
		t.Errorf("expected (%v, true), got (%v, %t)", want, result, ok)
	}
}

func TestIntersection_RepeatedElements(t *testing.T) {
	a := []int{65, 3, 3, 58, 678, 64}
	b := []int{64, 2, 3, 3, 43}
	want := []int{3, 3, 64}

	ok, result := Intersection(a, b)

	if !slices.Equal(result, want) || !ok {
		t.Errorf("expected (%v, true), got (%v, %t)", want, result, ok)
	}
}

func TestIntersection_EmptySlice(t *testing.T) {
	a := []int{}
	b := []int{64, 2, 3, 3, 43}
	want := []int{}

	ok, result := Intersection(a, b)

	if !slices.Equal(result, want) || ok {
		t.Errorf("expected (%v, false), got (%v, %t)", want, result, ok)
	}
}

func TestIntersection_NilSlice(t *testing.T) {
	var a []int
	var b []int
	want := []int{}

	ok, result := Intersection(a, b)

	if !slices.Equal(result, want) || ok {
		t.Errorf("expected (%v, false), got (%v, %t)", want, result, ok)
	}
}
