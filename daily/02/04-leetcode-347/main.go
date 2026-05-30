package main

import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n] += 1
	}
	reverseCountMap := make(map[int][]int)
	fs := []int{}
	for n, frequency := range countMap {
		cm, ok := reverseCountMap[frequency]
		if !ok {
			fs = append(fs, frequency)
		}
		reverseCountMap[frequency] = append(cm, n)
	}
	slices.SortFunc(fs, func(a, b int) int {
		return b - a
	})
	taken := 0
	res := []int{}
	for _, frequency := range fs {
		ns := reverseCountMap[frequency]
		for _, n := range ns {
			res = append(res, n)
			taken += 1
			if taken >= k {
				break
			}
		}
		if taken >= k {
			break
		}
	}

	return res
}
