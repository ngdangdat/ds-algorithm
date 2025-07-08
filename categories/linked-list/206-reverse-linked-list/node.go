package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) IsEqual(head *ListNode) bool {
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

func (n *ListNode) Print() {
	for n.Next != nil {
		fmt.Printf("%d ", n.Val)
		n = n.Next
	}
	fmt.Printf("%d ", n.Val)
	fmt.Println("")
}
