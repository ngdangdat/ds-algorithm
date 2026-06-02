package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	const TARGET = 0
	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if (l > i+1) && nums[l] == nums[l-1] {
				l++
				continue
			}
			if (r < len(nums)-1) && nums[r] == nums[r+1] {
				r--
				continue
			}
			sumILR := nums[i] + nums[l] + nums[r]
			if sumILR > TARGET {
				r--
				continue
			}
			if sumILR == TARGET {
				result = append(result, []int{nums[i], nums[l], nums[r]})
			}
			l++
		}
	}
	return result
}
