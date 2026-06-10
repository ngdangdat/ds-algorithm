package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, 1
	for r < len(nums) {
		for r < len(nums)-1 && nums[r] == nums[l] {
			r++
		}
		if r < len(nums) && nums[r] != nums[l] {
			if l < r-1 {
				// swap needed here
				l++
				nums[l] = nums[r]
			} else if l == r-1 {
				l++
			}
		}
		r++
	}
	return l + 1
}
