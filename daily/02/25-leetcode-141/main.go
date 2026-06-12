package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	f, s := head, head
	i := 1
	for f != nil && s != nil {
		f = f.Next
		if i%2 == 0 {
			s = s.Next
		}
		if f == s {
			return true
		}
		i++
	}
	return false
}
