package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	l1, l2 := m-1, n-1
	r := m + n - 1
	for r >= 0 && l1 >= 0 && l2 >= 0 {
		if nums1[l1] > nums2[l2] {
			nums1[r] = nums1[l1]
			l1--
		} else {
			nums1[r] = nums2[l2]
			l2--
		}
		r--
	}
	for r >= 0 && l1 >= 0 {
		nums1[r] = nums1[l1]
		l1--
		r--
	}
	for r >= 0 && l2 >= 0 {
		nums1[r] = nums2[l2]
		l2--
		r--
	}
}
