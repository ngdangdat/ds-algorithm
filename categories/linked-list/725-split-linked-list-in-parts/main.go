package main

import "fmt"

func getNodeLength(head *ListNode) int {
	tmp := head
	nodeLength := 0
	for tmp != nil {
		nodeLength += 1
		tmp = tmp.Next
	}
	return nodeLength
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var res []*ListNode
	n := getNodeLength(head)
	size, remainder := n/k, n%k
	for k > 0 {
		partSize := size
		if remainder > 0 {
			partSize += 1
			remainder -= 1
		}
		var prev, next *ListNode = head, nil
		for partSize > 1 && prev != nil {
			partSize -= 1
			prev = prev.Next
		}
		if prev != nil {
			next = prev.Next
			prev.Next = nil
		}
		res = append(res, head)
		head = next
		k -= 1
	}
	return res
}

func main() {
	ints := []int{1, 2, 3}
	k := 5
	lInts := ListNodeFromIntList(ints)
	res := splitListToParts(lInts, k)
	fmt.Printf("%v\n", res)
}
