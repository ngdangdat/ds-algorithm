package main

import "fmt"

func findMid(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right {
		fast = fast.Next
		if fast != right {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return slow
}

func merge(left, right *ListNode) *ListNode {
	head := &ListNode{}
	curr := head

	for left != nil && right != nil {
		if left.Val > right.Val {
			curr.Next = right
			right = right.Next
		} else {
			curr.Next = left
			left = left.Next
		}
		curr = curr.Next
	}
	if left != nil {
		curr.Next = left
	}
	if right != nil {
		curr.Next = right
	}
	return head.Next
}

func mergeSort(left, right *ListNode) *ListNode {
	if left == right {
		return left
	}
	if left.Next == right {
		if right == nil {
			return left
		}
		if left.Val > right.Val {
			left.Next = nil
			return merge(right, left)
		}
		return left
	}
	mid := findMid(left, right)
	nextMid := mid.Next
	mid.Next = nil
	// fmt.Printf("left: %v, right: %v, mid: %v\n", left, right, mid)
	mergedLeft := mergeSort(left, mid)
	mergedRight := mergeSort(nextMid, right)
	// fmt.Printf("mergedLeft: %v, mergedRight: %v, mid: %v\n", mergedLeft, mergedRight, mid)
	return merge(mergedLeft, mergedRight)
}

// sortList sorts a linked list in O(n log n) time and O(1) space
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result := mergeSort(head, nil)
	return result
}

func main() {
	// Test case example
	fmt.Println("148. Sort List")
	// Test case: [4,2,1,3] -> [1,2,3,4]
	head := ListNodeFromIntList([]int{4, 2, 1, 3})
	fmt.Print("Input: ")
	head.Print()
	sorted := sortList(head)
	fmt.Print("Output: ")
	sorted.Print()
}
