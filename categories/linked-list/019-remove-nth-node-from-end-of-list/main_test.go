package main

import (
	"fmt"
	"testing"
)

func TestRemoveNthNodeFromEndOfLis(t *testing.T) {
	cases := []struct {
		InputList []int
		N         int
		Expected  []int
	}{
		{
			InputList: []int{1, 2, 3, 4, 5},
			N:         2,
			Expected:  []int{1, 2, 3, 5},
		},
		{
			InputList: []int{1, 2},
			N:         1,
			Expected:  []int{1},
		},
		{
			InputList: []int{1},
			N:         1,
			Expected:  []int{},
		},
		{
			InputList: []int{1, 2, 3, 4, 5},
			N:         5,
			Expected:  []int{2, 3, 4, 5},
		},
		{
			InputList: []int{1, 2, 3, 4, 5, 6, 7},
			N:         1,
			Expected:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			InputList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			N:         4,
			Expected:  []int{1, 2, 3, 4, 5, 6, 8, 9, 10},
		},
		{
			InputList: []int{1, 2},
			N:         2,
			Expected:  []int{2},
		},
		{
			InputList: []int{1, 2, 3},
			N:         3,
			Expected:  []int{2, 3},
		},
		{
			InputList: []int{1, 2, 3},
			N:         2,
			Expected:  []int{1, 3},
		},
		{
			InputList: []int{1, 2, 3, 4},
			N:         2,
			Expected:  []int{1, 2, 4},
		},
	}

	for index, tCase := range cases {
		input := ListNodeFromIntList(tCase.InputList)
		got := removeNthFromEnd(input, tCase.N)
		fmt.Printf("input: %v, got: %v\n", input, got)
		expected := ListNodeFromIntList(tCase.Expected)
		if !got.IsEqual(expected) {
			t.Errorf("Case %d failed, expected: %v, got: %v", index, expected, got)
		}
	}
}
