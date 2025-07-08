package main

import "testing"

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		List1    []int
		List2    []int
		Expected []int
	}{
		{
			List1:    []int{1, 2, 3},
			List2:    []int{1, 4, 5},
			Expected: []int{1, 1, 2, 3, 4, 5},
		},
		{
			List1:    []int{1, 1, 1, 1, 1},
			List2:    []int{1, 1, 1, 1, 2},
			Expected: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		},
		{
			List1:    []int{1, 2, 4},
			List2:    []int{1, 3, 4},
			Expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			List1:    []int{5},
			List2:    []int{1, 2, 4},
			Expected: []int{1, 2, 4, 5},
		},
		{
			List1:    []int{},
			List2:    []int{},
			Expected: []int{},
		},
	}
	for _, tCase := range cases {
		input1, input2 := ListNodeFromIntList(tCase.List1), ListNodeFromIntList(tCase.List2)
		result := mergeTwoLists(input1, input2)
		expected := ListNodeFromIntList(tCase.Expected)
		if !result.IsEqual(expected) {
			t.Errorf("invalid case, expected: %v, got: %v", expected, result)
		}
	}
}
