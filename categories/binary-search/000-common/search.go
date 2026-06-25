package main

import "fmt"

func binSearch(nums []int, target int, left int, right int) int {
	if left < right {
		mid := left + (right-left+1)/2
		fmt.Printf("left=%d, right=%d, mid=%d\n", left, right, mid)
		fmt.Printf("nums[left]=%d, nums[right]=%d, nums[mid]=%d\n", nums[left], nums[right], nums[mid])
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return binSearch(nums, target, left, mid-1)
		} else {
			return binSearch(nums, target, mid+1, right)
		}
	}
	return -1
}

func search(nums []int, target int) int {
	fmt.Printf("nums=%v\n", nums)
	return binSearch(nums, target, 0, len(nums)-1)
}

func main() {
	nums1 := []int{1, 4, 5, 7, 9, 21}

	res1 := search(nums1, 7)
	fmt.Printf("7 res1=%d\n", res1)
	res2 := search(nums1, 78)
	fmt.Printf("78 res2=%d\n", res2)

}
