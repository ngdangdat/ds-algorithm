package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := range len(nums) - 3 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < (len(nums) - 2); j++ {
			if j > (i+1) && nums[j-1] == nums[j] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				if l > (j+1) && nums[l] == nums[l-1] {
					l++
					continue
				}
				if r < (len(nums)-1) && nums[r] == nums[r+1] {
					r--
					continue
				}
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
					continue
				}
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
				}
				l++
			}
		}
	}
	return result
}
