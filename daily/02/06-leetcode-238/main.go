package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	leftMul := 1
	for i := range len(nums) {
		if i == 0 {
			continue
		}
		leftMul *= nums[i-1]
		res[i] = leftMul
	}
	rightMul := 1
	for r := len(nums) - 1; r >= 0; r-- {
		if r == len(nums)-1 {
			continue
		}
		rightMul *= nums[r+1]
		mul := rightMul
		if r != 0 {
			mul *= res[r]
		}
		res[r] = mul
	}

	return res
}
