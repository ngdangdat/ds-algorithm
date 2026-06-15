package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"example1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"example2", []int{1, 1}, 1},
		{"two_tall", []int{4, 3, 2, 1, 4}, 16},
		{"descending", []int{5, 4, 3, 2, 1}, 6},
		{"ascending", []int{1, 2, 3, 4, 5}, 6},
		{"all_equal", []int{3, 3, 3, 3}, 9},
		{"single", []int{5}, 0},
		{"empty", []int{}, 0},
		{"with_zeros", []int{0, 2, 0, 4, 0}, 4},
		{"peak_middle", []int{1, 2, 4, 3}, 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxArea(tc.height)
			if got != tc.want {
				t.Errorf("maxArea(%v) = %d, want %d", tc.height, got, tc.want)
			}
		})
	}
}

// TestMaxAreaPreservesInput guards that the input slice is not mutated.
func TestMaxAreaPreservesInput(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	snapshot := append([]int(nil), height...)
	maxArea(height)
	for i := range height {
		if height[i] != snapshot[i] {
			t.Errorf("input mutated at index %d: got %v, want %v", i, height, snapshot)
		}
	}
}
