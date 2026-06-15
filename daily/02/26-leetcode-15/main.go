package main

import (
	"fmt"
	"sort"
)

// threeSum returns all unique triplets [nums[i], nums[j], nums[k]] such that
// i != j != k and nums[i] + nums[j] + nums[k] == 0.

func threeSum(nums []int) [][]int {
	const TARGET = 0
	res := [][]int{}
	existMap := make(map[string]bool)
	processed := make(map[int]bool)
	for i := range len(nums) {
		if processed[nums[i]] {
			continue
		}
		c := TARGET - nums[i]
		needs := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			// we start with two sum here
			need := c - nums[j]
			if needs[need] {
				r := []int{nums[i], nums[j], need}
				sort.Ints(r)
				k := fmt.Sprintf("%d_%d_%d", r[0], r[1], r[2])
				if !existMap[k] {
					existMap[k] = true
					res = append(res, r)
				}
			} else {
				needs[nums[j]] = true
			}
		}
		processed[nums[i]] = true
	}
	return res
}
