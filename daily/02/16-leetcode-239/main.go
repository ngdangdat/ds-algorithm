package main

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	deque := []int{}
	for i, num := range nums {
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
