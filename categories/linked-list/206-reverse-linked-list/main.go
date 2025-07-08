package main

import "fmt"



func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev, next *ListNode = nil, head
	for head != nil {
		next = head.Next
		// move the next to back node
		head.Next = prev
		// set prev
		prev = head
		// set next
		head = next
	}
	return prev
}

func reverse(prev *ListNode, head *ListNode) *ListNode {
	if head == nil {
		return prev
	}
	var next *ListNode = head.Next
	head.Next = prev
	return reverse(head, next)
}

func reverseListRecursion(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func wrongReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newTail *ListNode = nil
	for head.Next != nil {
		node := &ListNode{
			Val:  head.Val,
			Next: newTail,
		}
		newTail = node
		head = head.Next
	}
	newTail = &ListNode{
		Val:  head.Val,
		Next: newTail,
	}
	return newTail
}

func main() {
	// head = [1, 2]
	// expected = [2, 1]
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	head.Print()
	rev := reverseList(head)
	rev.Print()
	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	revRecur := reverseListRecursion(head)
	revRecur.Print()
}
