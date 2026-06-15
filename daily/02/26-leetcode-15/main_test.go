package main

import (
	"reflect"
	"sort"
	"testing"
)

// normalize sorts each triplet and then sorts the slice of triplets so that
// results can be compared regardless of ordering.
func normalize(triplets [][]int) [][]int {
	out := make([][]int, len(triplets))
	for i, t := range triplets {
		cp := append([]int(nil), t...)
		sort.Ints(cp)
		out[i] = cp
	}
	sort.Slice(out, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return false
	})
	return out
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"example1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"all_zeros", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"no_triplet", []int{0, 1, 1}, [][]int{}},
		{"empty", []int{}, [][]int{}},
		{"too_short", []int{1, -1}, [][]int{}},
		{"duplicates", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"all_positive", []int{1, 2, 3, 4}, [][]int{}},
		{"all_negative", []int{-1, -2, -3, -4}, [][]int{}},
		{"multiple", []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			[][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
		{"four_zeros", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := normalize(threeSum(tc.nums))
			want := normalize(tc.want)
			if len(got) == 0 && len(want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("threeSum(%v) = %v, want %v", tc.nums, got, want)
			}
		})
	}
}

// TestThreeSumNoDuplicates guards that the result contains no duplicate triplets.
func TestThreeSumNoDuplicates(t *testing.T) {
	got := normalize(threeSum([]int{-1, 0, 1, 2, -1, -4, -1, 0, 1}))
	seen := map[[3]int]bool{}
	for _, tr := range got {
		key := [3]int{tr[0], tr[1], tr[2]}
		if seen[key] {
			t.Errorf("duplicate triplet found: %v", tr)
		}
		seen[key] = true
	}
}
