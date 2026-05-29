package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	numIndexMap := make(map[int]int)
	for idx, n := range nums {
		complement := target - n
		index, ok := numIndexMap[complement]
		if !ok {
			numIndexMap[n] = idx
		}
		if ok {
			res[0] = index
			res[1] = idx
			break
		}
	}
	return res
}
