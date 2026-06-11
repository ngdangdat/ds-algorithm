package main

func moveZeroes(nums []int) {
	l, r := 0, 1
	for r < len(nums) {
		for l < len(nums) && nums[l] != 0 {
			l++
		}
		if l < r && nums[r] != 0 {
			nums[l] = nums[r]
			nums[r] = 0
			l++
		}
		r++
	}
}
