package main

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{
		Next: head,
	}
	prev := res
	index := 0
	for head != nil {
		if index >= n {
			prev = prev.Next
		}
		index += 1
		head = head.Next
	}
	if prev.Next != nil {
		prev.Next = prev.Next.Next
	}

	return res.Next
}

func main() {
	input := ListNodeFromIntList([]int{1, 2, 3})
	inputN := 2
	output := removeNthFromEnd(input, inputN)
	fmt.Printf("value: %v", output)
}
