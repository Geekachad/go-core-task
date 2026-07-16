package main

import (
	"maps"
	"slices"
	"testing"
)

func TestSliceToSet(t *testing.T) {
	slice := []string{"banana", "date", "fig"}
	want := map[string]struct{}{
		"banana": {},
		"date":   {},
		"fig":    {},
	}

	result := SliceToSet(slice)

	if !maps.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestSliceToSet_EmptySlice(t *testing.T) {
	slice := []string{}
	want := map[string]struct{}{}

	result := SliceToSet(slice)

	if !maps.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestSliceToSet_NilSlice(t *testing.T) {
	var slice []string
	want := map[string]struct{}{}

	result := SliceToSet(slice)

	if !maps.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestPresentRightAbsentLeft(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	want := []string{"apple", "cherry", "43", "lead", "gno1"}

	result := PresentRightAbsentLeft(slice1, slice2)

	if !slices.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestPresentRightAbsentLeft_RepeatedElements(t *testing.T) {
	slice1 := []string{"apple", "apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "banana", "date", "fig"}
	want := []string{"apple", "apple", "cherry", "43", "lead", "gno1"}

	result := PresentRightAbsentLeft(slice1, slice2)

	if !slices.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestPresentRightAbsentLeft_EmptySlice(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{}
	want := []string{}

	result := PresentRightAbsentLeft(slice1, slice2)

	if !slices.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func TestPresentRightAbsentLeft_NilSlice(t *testing.T) {
	var slice1 []string
	var slice2 []string
	want := []string{}

	result := PresentRightAbsentLeft(slice1, slice2)

	if !slices.Equal(result, want) {
		t.Errorf("expected %v, got %v", want, result)
	}
}
