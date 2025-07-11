package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeFromIntList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	
	head := &ListNode{Val: values[0]}
	current := head
	
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	
	return head
}

func (head *ListNode) String() string {
	if head == nil {
		return "nil"
	}
	
	result := fmt.Sprintf("%d", head.Val)
	current := head.Next
	
	for current != nil {
		result += fmt.Sprintf(" -> %d", current.Val)
		current = current.Next
	}
	
	return result
}