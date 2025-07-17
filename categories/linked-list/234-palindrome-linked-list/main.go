package main

import "fmt"

func findMid(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return slow
}

// isPalindrome checks if a linked list is a palindrome
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := findMid(head)
	// fmt.Printf("mid: %v\n", mid)
	prev, curr := mid, mid.Next
	prev.Next = nil
	for curr != nil {
		next := curr.Next
		// fmt.Printf("prev: %v, curr: %v, next: %v\n", prev, curr, next)
		curr.Next = prev

		prev = curr
		curr = next
		// fmt.Printf("prev: %v, curr: %v\n", prev, curr)
	}
	reverseHead := prev
	isEqual := true
	for head != mid {
		if head.Val != reverseHead.Val {
			isEqual = false
			break
		}

		reverseHead = reverseHead.Next
		head = head.Next
	}
	return isEqual
}

func main() {
	fmt.Println("234. Palindrome Linked List")

	// Test case 1: [1,2,2,1] - should return true
	head1 := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}}
	fmt.Printf("Test 1 [1,2,2,2,1]: %t\n", isPalindrome(head1))

	// Test case 2: [1,2] - should return false
	head2 := &ListNode{1, &ListNode{2, nil}}
	fmt.Printf("Test 2 [1,2]: %t\n", isPalindrome(head2))

	// Test case 3: [1] - should return true
	head3 := &ListNode{1, nil}
	fmt.Printf("Test 3 [1]: %t\n", isPalindrome(head3))

	// Test case 4: [1,2,3,2,1] - should return true
	head4 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	fmt.Printf("Test 4 [1,2,3,2,1]: %t\n", isPalindrome(head4))

	// Test case 5: empty list - should return true
	fmt.Printf("Test 5 []: %t\n", isPalindrome(nil))
}
