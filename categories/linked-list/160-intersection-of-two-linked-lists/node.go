package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) IsEqual(head *ListNode) bool {
	if n == nil {
		return head == nil
	}
	for n.Next != nil && head.Next != nil {
		if n.Val != head.Val {
			return false
		}
		n = n.Next
		head = head.Next
	}
	if n.Val != head.Val {
		return false
	}
	if (n.Next == nil && head.Next != nil) || (n.Next != nil && head.Next == nil) {
		return false
	}
	return true
}

func (n *ListNode) AddNodes(head *ListNode) {
	c := n
	for c.Next != nil {
		c = c.Next
	}
	c.Next = head
}

func (n *ListNode) Print() {
	for n.Next != nil {
		fmt.Printf("%d ", n.Val)
		n = n.Next
	}
	fmt.Printf("%d ", n.Val)
	fmt.Println("")
}

func (n ListNode) String() string {
	head := &n
	ns := []int{}
	for head != nil {
		ns = append(ns, head.Val)
		head = head.Next
	}
	return fmt.Sprintf("{%p}ListNode:%v", &n, ns)
}

func ListNodeFromIntList(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}

	if len(l) == 1 {
		return &ListNode{
			Val:  l[0],
			Next: nil,
		}
	}

	head := &ListNode{
		Val:  l[0],
		Next: nil,
	}
	curr := head

	for i, val := range l {
		if i == 0 {
			continue
		}

		curr.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		curr = curr.Next
	}

	return head
}
