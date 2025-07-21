package main

import (
	"fmt"
	"testing"
)

func TestKokoEatingBanana(t *testing.T) {
	cases := []struct {
		piles    []int
		hours    int
		expected int
	}{
		{
			piles:    []int{3, 6, 7, 11},
			hours:    8,
			expected: 4,
		},
		{
			piles:    []int{30, 11, 23, 4, 20},
			hours:    5,
			expected: 30,
		},
		{
			piles:    []int{30, 11, 23, 4, 20},
			hours:    6,
			expected: 23,
		},
		{
			piles:    []int{1, 1, 1, 999999999},
			hours:    10,
			expected: 142857143,
		},
		{
			piles:    []int{312884470},
			hours:    312884469,
			expected: 2,
		},
		{
			piles:    []int{1},
			hours:    1,
			expected: 1,
		},
		{
			piles:    []int{1, 1, 1, 1},
			hours:    8,
			expected: 1,
		},
		{
			piles:    []int{5, 5, 5, 5, 5},
			hours:    10,
			expected: 3,
		},
	}

	for index, testCase := range cases {
		caseName := fmt.Sprintf("Normal case %d", index)
		t.Run(caseName, func(t *testing.T) {
			got := minEatingSpeed(testCase.piles, testCase.hours)
			if got != testCase.expected {
				t.Errorf("mismatch, expected: %d, got: %d", testCase.expected, got)
			}
		})
	}
}
