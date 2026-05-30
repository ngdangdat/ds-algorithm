package main

import (
	"slices"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	cases := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			// No zeros: matrix must be left unchanged.
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{
			// Single zero in the first row/col corner: exercises the
			// first-row / first-col marker overlap.
			matrix: [][]int{
				{0, 2},
				{3, 4},
			},
			expected: [][]int{
				{0, 0},
				{0, 4},
			},
		},
	}
	for index, c := range cases {
		setZeroes(c.matrix)
		for r := range c.expected {
			if !slices.Equal(c.matrix[r], c.expected[r]) {
				t.Fatalf("Case %d row %d fails, expected=%v got=%v", index, r, c.expected[r], c.matrix[r])
			}
		}
	}
}
