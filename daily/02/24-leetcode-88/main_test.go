package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{"example1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"example2", []int{1}, 1, []int{}, 0, []int{1}},
		{"example3", []int{0}, 0, []int{1}, 1, []int{1}},
		{"nums2 empty", []int{1, 2, 3}, 3, []int{}, 0, []int{1, 2, 3}},
		{"nums1 empty", []int{0, 0, 0}, 0, []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"interleaved", []int{1, 3, 5, 0, 0, 0}, 3, []int{2, 4, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"all nums2 smaller", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"all nums2 larger", []int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"duplicates across", []int{1, 2, 2, 0, 0}, 3, []int{2, 2}, 2, []int{1, 2, 2, 2, 2}},
		{"negatives", []int{-5, -2, 0, 0, 0}, 2, []int{-3, -1, 4}, 3, []int{-5, -3, -2, -1, 4}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			if !reflect.DeepEqual(tc.nums1, tc.want) {
				t.Errorf("merge(%v, %d, %v, %d) => %v, want %v",
					tc.nums1, tc.m, tc.nums2, tc.n, tc.nums1, tc.want)
			}
		})
	}
}
