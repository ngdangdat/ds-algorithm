package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for i, c := range cases {
		got := twoSum(c.nums, c.target)
		if !reflect.DeepEqual(got, c.expected) {
			t.Fatalf("failed case %d, expected=%v got=%v", i+1, c.expected, got)
		}
	}
}
