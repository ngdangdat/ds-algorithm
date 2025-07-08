package main

import "fmt"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}
	// dummy node
	start := &ListNode{
		Val:  -1,
		Next: nil,
	}
	curr := start
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			curr.Next = list2
			list2 = list2.Next
		} else {
			curr.Next = list1
			list1 = list1.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return start.Next
}

func main() {
	l1 := ListNodeFromIntList([]int{-10, -10, -9, -4, 1, 6, 6})
	l2 := ListNodeFromIntList([]int{-7})
	fmt.Printf("l1: %v, l2: %v\n", l1, l2)
	result := mergeTwoLists(l1, l2)
	fmt.Printf("l1: %v, l2: %v, result: %v\n", l1, l2, result)

	l1 = ListNodeFromIntList([]int{1, 2, 3})
	l2 = ListNodeFromIntList([]int{1, 4, 5})
	fmt.Printf("l1: %v, l2: %v\n", l1, l2)
	result = mergeTwoLists(l1, l2)
	fmt.Printf("l1: %v, l2: %v, result: %v\n", l1, l2, result)
}
