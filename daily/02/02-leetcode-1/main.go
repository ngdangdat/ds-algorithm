package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	numIndexMap := make(map[int]int)
	for idx, n := range nums {
		complement := target - n
		index, ok := numIndexMap[complement]
		if ok {
			res[0] = index
			res[1] = idx
			break
		} else {
			numIndexMap[n] = idx
		}
	}
	return res
}
