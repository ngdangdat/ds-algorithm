package main

func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]bool)
	for _, n := range nums1 {
		nums1Map[n] = true
	}
	res := []int{}
	existMap := make(map[int]bool)
	for _, n := range nums2 {
		if !existMap[n] && nums1Map[n] {
			existMap[n] = true
			res = append(res, n)
		}
	}
	return res
}
