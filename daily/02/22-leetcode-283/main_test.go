package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single zero", []int{0}, []int{0}},
		{"single non-zero", []int{1}, []int{1}},
		{"no zeroes", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all zeroes", []int{0, 0, 0}, []int{0, 0, 0}},
		{"example1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"leading zeroes", []int{0, 0, 1}, []int{1, 0, 0}},
		{"trailing zeroes", []int{1, 0, 0}, []int{1, 0, 0}},
		{"order preserved", []int{4, 0, 5, 0, 0, 3, 6}, []int{4, 5, 3, 6, 0, 0, 0}},
		{"negatives", []int{0, -1, 0, -2}, []int{-1, -2, 0, 0}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			moveZeroes(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.want) {
				t.Errorf("moveZeroes() = %v, want %v", tc.nums, tc.want)
			}
		})
	}
}
