package main

import "testing"

// buildList builds a linked list from vals. If pos >= 0, the tail's Next is
// wired back to the node at index pos to form a cycle (LeetCode's convention).
// pos == -1 means no cycle.
func buildList(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i < len(vals)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if pos >= 0 {
		nodes[len(vals)-1].Next = nodes[pos]
	}
	return nodes[0]
}

func TestHasCycle(t *testing.T) {
	cases := []struct {
		name string
		vals []int
		pos  int // index tail links to; -1 = no cycle
		want bool
	}{
		{"empty list", []int{}, -1, false},
		{"single no cycle", []int{1}, -1, false},
		{"single self cycle", []int{1}, 0, true},
		{"two no cycle", []int{1, 2}, -1, false},
		{"two cycle to head", []int{1, 2}, 0, true},
		{"example1", []int{3, 2, 0, -4}, 1, true},
		{"long no cycle", []int{1, 2, 3, 4, 5}, -1, false},
		{"cycle to middle", []int{1, 2, 3, 4, 5}, 2, true},
		{"cycle tail to itself", []int{1, 2, 3}, 2, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			head := buildList(tc.vals, tc.pos)
			if got := hasCycle(head); got != tc.want {
				t.Errorf("hasCycle(%v, pos=%d) = %v, want %v", tc.vals, tc.pos, got, tc.want)
			}
		})
	}
}
