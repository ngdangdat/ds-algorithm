package main

func productExceptSelf(nums []int) []int {
	zeroCount := 0
	res := make([]int, len(nums))
	mul := 1
	for _, n := range nums {
		if n == 0 {
			zeroCount += 1
			if zeroCount >= 2 {
				return res
			}
			continue
		}
		mul *= n
	}
	for i, n := range nums {
		if n == 0 {
			res[i] = mul
			continue
		}
		if zeroCount == 1 {
			res[i] = 0
			continue
		}
		res[i] = mul / nums[i]
	}
	return res
}
