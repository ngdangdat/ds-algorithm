package main

import "testing"

type Input struct {
	ListA    []int
	ListB    []int
	Expected []int
}

func TestIntersectionNode(t *testing.T) {
	cases := []Input{
		{
			ListA:    []int{1, 4, 9},
			ListB:    []int{2, 5, 10},
			Expected: []int{11, 29, 30},
		},
		{
			ListA:    []int{4, 1},
			ListB:    []int{5, 6, 1},
			Expected: []int{8, 4, 5},
		},
		{
			ListA:    []int{1, 9, 1},
			ListB:    []int{3},
			Expected: []int{2, 4},
		},
		{
			ListA:    []int{2, 6, 4},
			ListB:    []int{1, 5},
			Expected: []int{},
		},
	}

	for _, tCase := range cases {
		lA := ListNodeFromIntList(tCase.ListA)
		lB := ListNodeFromIntList(tCase.ListB)
		expected := ListNodeFromIntList(tCase.Expected)
		lA.AddNodes(expected)
		lB.AddNodes(expected)
		got := getIntersectionNode(lA, lB)
		if expected != nil {
			if got != expected {
				t.Errorf("lA: %v\nlB: %v\nfailed case, expected: %v, got: %v", lA, lB, expected, got)
			}
			continue
		}
		if got != nil {
			t.Errorf("lA: %v\nlB: %v\nfailed case, expected: %v, got: %v", lA, lB, expected, got)
		}
	}
}
