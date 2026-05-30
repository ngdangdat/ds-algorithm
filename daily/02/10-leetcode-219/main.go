package main

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)
	for index, n := range nums {
		latestIndex, ok := indexMap[n]
		if ok {
			if (index - latestIndex) <= k {
				return true
			}
		}
		indexMap[n] = index
	}
	return false
}
