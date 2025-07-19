package main

import "fmt"

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
		if fast == slow {
			return true
		}
	}

	return false
}

func main() {
	headex1 := &ListNode{Val: 1}
	ex1 := &ListNode{Val: 1, Next: &ListNode{Val: -1, Next: headex1}}
	headex1.Next = ex1
	result1 := hasCycle(headex1)
	fmt.Printf("Case 1, expect: true, result: %v\n", result1)

	headex2 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	fmt.Printf("Case 2, expect: false, result: %v", hasCycle(headex2))
}
