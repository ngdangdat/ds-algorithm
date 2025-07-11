package main

import "fmt"

/*
Problem: Add Two Numbers
URL: https://leetcode.com/problems/add-two-numbers/

Description:
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:
- The number of nodes in each linked list is in the range [1, 100].
- 0 <= Node.val <= 9
- It is guaranteed that the list represents a number that does not have leading zeros.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	remainder := 0
	res := &ListNode{}
	headRes := res
	for l1 != nil || l2 != nil || remainder != 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		sum := val1 + val2 + remainder
		res.Next = &ListNode{Val: (sum % 10)}
		remainder = sum / 10
		res = res.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return headRes.Next
}

func main() {
	// Test case 1: [2,4,3] + [5,6,4] = [7,0,8]
	l1 := ListNodeFromIntList([]int{2, 4, 3})
	l2 := ListNodeFromIntList([]int{5, 6, 4})
	result := addTwoNumbers(l1, l2)
	fmt.Printf("Test 1 - Input: %v + %v\n", l1, l2)
	fmt.Printf("Test 1 - Output: %v\n\n", result)

	// Test case 2: [0] + [0] = [0]
	l1 = ListNodeFromIntList([]int{0})
	l2 = ListNodeFromIntList([]int{0})
	result = addTwoNumbers(l1, l2)
	fmt.Printf("Test 2 - Input: %v + %v\n", l1, l2)
	fmt.Printf("Test 2 - Output: %v\n\n", result)

	// Test case 3: [9,9,9,9,9,9,9] + [9,9,9,9] = [8,9,9,9,0,0,0,1]
	l1 = ListNodeFromIntList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = ListNodeFromIntList([]int{9, 9, 9, 9})
	result = addTwoNumbers(l1, l2)
	fmt.Printf("Test 3 - Input: %v + %v\n", l1, l2)
	fmt.Printf("Test 3 - Output: %v\n", result)
}
