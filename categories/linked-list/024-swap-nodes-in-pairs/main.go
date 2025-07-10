package main

import "fmt"

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{Next: head}
	prev := res
	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next

		second := first.Next
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
	}

	return res.Next
}

func main() {
	// Example: [1,2,3,4] -> [2,1,4,3]
	head := ListNodeFromIntList([]int{1, 2, 3, 4})

	fmt.Println("Original:")
	head.Print()

	result := swapPairs(head)

	fmt.Println("After swapping pairs:")
	result.Print()
}
