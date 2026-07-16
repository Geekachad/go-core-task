package main

import (
	"slices"
	"testing"
)

func TestSliceExample(t *testing.T) {
	type testCase struct {
		name  string // Понятное имя теста для вывода в консоли
		input []int  // Входные данные
		want  []int  // Ожидаемый результат
	}

	testCases := []testCase{
		{
			name:  "slice example",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{2, 4},
		},
		{
			name:  "no even numbers",
			input: []int{1, 3, 5},
			want:  []int{},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  []int{},
		},
		{
			name:  "repeating numbers",
			input: []int{2, 2, 3, 4, 5},
			want:  []int{2, 2, 4},
		},
		{
			name:  "zero value",
			input: []int{0},
			want:  []int{0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sliceExample(tc.input)

			if !slices.Equal(result, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, result)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	type testCase struct {
		name  string // Понятное имя теста для вывода в консоли
		input []int  // Входные данные
		want  []int  // Ожидаемый результат
	}

	testCases := []testCase{
		{
			name:  "add elements",
			input: []int{1, 2},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{3},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  []int{3},
		},
		{
			name:  "repeating numbers",
			input: []int{3},
			want:  []int{3, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := addElements(tc.input, 3)

			if !slices.Equal(result, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, result)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	type testCase struct {
		name  string
		input []int
	}

	testCases := []testCase{
		{
			name:  "copy slice",
			input: []int{1, 2, 3},
		},
		{
			name:  "empty slice",
			input: []int{},
		},
		{
			name:  "nil slice",
			input: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := copySlice(tc.input)
			tc.input = append(tc.input, 4)

			if slices.Equal(result, tc.input) {
				t.Errorf("expected %v, got %v", tc.input, result)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	type testCase struct {
		name  string
		idx   int
		input []int
		want  []int
	}

	testCases := []testCase{
		{
			name:  "remove element",
			idx:   0,
			input: []int{1, 2, 3},
			want:  []int{2, 3},
		},
		{
			name:  "remove last element",
			idx:   2,
			input: []int{1, 2, 3},
			want:  []int{1, 2},
		},
		{
			name:  "remove middle element",
			idx:   1,
			input: []int{1, 2, 3},
			want:  []int{1, 3},
		},
		{
			name:  "remove non-existent element",
			idx:   2,
			input: []int{1, 2},
			want:  nil,
		},
		{
			name:  "wrong index",
			idx:   -1,
			input: []int{1, 2},
			want:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeElement(tc.input, tc.idx)

			if !slices.Equal(result, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, result)
			}
		})
	}
}
