package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	res := []int{}
	for l < r {
		s := numbers[l] + numbers[r]
		if s < target {
			l++
		} else if s > target {
			r--
		} else {
			res = append(res, l+1)
			res = append(res, r+1)
			break
		}
	}
	return res
}
