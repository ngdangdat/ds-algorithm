package main

func containsDuplicate(nums []int) bool {
	numCountMap := make(map[int]int)
	for _, num := range nums {
		numCountMap[num] = numCountMap[num] + 1
		if numCountMap[num] >= 2 {
			return true
		}
	}
	return false
}
