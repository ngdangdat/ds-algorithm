package main

import "fmt"

// oddEvenList rearranges the linked list to group odd and even positioned nodes
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddNode := head
	evenNode := head.Next
	evenStart := evenNode
	prevNode := head

	for oddNode != nil && evenNode != nil {
		oddNode.Next = evenNode.Next
		prevNode = oddNode
		oddNode = oddNode.Next
		if oddNode != nil {
			evenNode.Next = oddNode.Next
			evenNode = evenNode.Next
		}
	}
	if oddNode != nil {
		oddNode.Next = evenStart
	} else {
		prevNode.Next = evenStart
	}
	return head
}

func main() {
	fmt.Println("328. Odd Even Linked List")

	// Test case 1: [1,2] -> [1,2]
	fmt.Println("\nTest case 1: 4 elements")
	nums1 := []int{1, 2, 3, 4}
	head1 := ListNodeFromIntList(nums1)

	fmt.Print("Original: ")
	head1.Print()

	result1 := oddEvenList(head1)
	fmt.Print("Result:   ")
	result1.Print()

	// Test case 2: [1,2,3,4,5] -> [1,3,5,2,4]
	fmt.Println("\nTest case 2: 5 elements")
	nums2 := []int{1, 2, 3, 4, 5}
	head2 := ListNodeFromIntList(nums2)

	fmt.Print("Original: ")
	head2.Print()

	result2 := oddEvenList(head2)
	fmt.Print("Result:   ")
	result2.Print()

	fmt.Println("\nTest case 3: 0 element")
	nums3 := []int{}
	head3 := ListNodeFromIntList(nums3)
	fmt.Print("Original: ")
	head3.Print()

	result3 := oddEvenList(head3)
	fmt.Print("Result:   ")
	result3.Print()
}
