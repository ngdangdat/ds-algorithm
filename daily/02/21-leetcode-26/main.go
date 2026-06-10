package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, 1
	for r < len(nums) {
		if nums[r] != nums[l] {
			nums[l+1] = nums[r]
			l++
		}
		r++
	}
	return l + 1
}
