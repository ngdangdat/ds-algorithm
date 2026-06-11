package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		name    string
		numbers []int
		target  int
		want    []int // 1-indexed [index1, index2]
	}{
		{"example1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"example2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"example3", []int{-1, 0}, -1, []int{1, 2}},
		{"two elements", []int{1, 2}, 3, []int{1, 2}},
		{"answer at ends", []int{1, 3, 5, 7, 9}, 10, []int{1, 5}},
		{"adjacent middle", []int{1, 2, 4, 7, 11}, 11, []int{3, 4}},
		{"with negatives", []int{-10, -3, 1, 4, 8}, -9, []int{1, 3}},
		{"duplicates as answer", []int{3, 3}, 6, []int{1, 2}},
		{"large gap", []int{0, 0, 3, 4}, 0, []int{1, 2}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := twoSum(tc.numbers, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tc.numbers, tc.target, got, tc.want)
			}
		})
	}
}
