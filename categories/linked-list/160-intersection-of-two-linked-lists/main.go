package main

func count(l *ListNode) int {
	if l == nil {
		return 0
	}
	res := 0
	h := l
	for h != nil {
		res += 1
		h = h.Next
	}
	return res
}

func getIntersection(large, small *ListNode, diff int) *ListNode {
	// fmt.Printf("diff: %d\n", diff)
	for diff != 0 {
		large = large.Next
		diff--
	}
	for large != small && large != nil && small != nil {
		// fmt.Printf("large: %p, small: %p\n", &large, &small)
		large = large.Next
		small = small.Next
	}

	return large
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	lA := count(headA)
	lB := count(headB)
	// fmt.Printf("lA: %d, lB: %d\n", lA, lB)

	if lA >= lB {
		return getIntersection(headA, headB, lA-lB)
	}

	return getIntersection(headB, headA, lB-lA)
}

func main() {

}
