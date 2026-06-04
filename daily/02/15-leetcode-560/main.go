package main

func subarraySum(nums []int, k int) int {
	res := 0
	prefixCountMap := make(map[int]int)
	prefixCountMap[0] = 1
	prefixSum := 0
	for _, num := range nums {
		prefixSum += num
		needed := prefixSum - k
		if prefixCountMap[needed] > 0 {
			res += prefixCountMap[needed]
		}
		prefixCountMap[prefixSum] += 1
	}
	return res
}
