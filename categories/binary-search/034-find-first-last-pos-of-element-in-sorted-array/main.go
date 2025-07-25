package main

import "fmt"

func searchMostRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
			continue
		} else if nums[mid] < target {
			l = mid + 1
			continue
		}

		if (mid < (len(nums) - 1)) && (nums[mid+1] == target) {
			l = mid + 1
			continue
		}

		res = mid
		break
	}

	return res
}

func searchMostLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
			continue
		} else if nums[mid] < target {
			l = mid + 1
			continue
		}

		if (mid != 0) && (nums[mid-1] == target) {
			r = mid - 1
			continue
		}

		res = mid
		break
	}

	return res
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	leftMostIndex := searchMostLeft(nums, target)
	if leftMostIndex == -1 {
		return res
	}

	rightMostIndex := searchMostRight(nums, target)
	res[0], res[1] = leftMostIndex, rightMostIndex

	return res
}

func main() {
	nums1 := []int{5, 7, 7, 8, 8, 10}
	target1 := 8
	expected1 := []int{3, 4}
	res1 := searchRange(nums1, target1)

	fmt.Printf("ex1, nums=%v, target=%d, expected=%v, got=%v\n", nums1, target1, expected1, res1)

	nums2 := []int{2, 2}
	target2 := 2
	expected2 := []int{0, 1}
	res2 := searchRange(nums2, target2)

	fmt.Printf("ex2, nums=%v, target=%d, expected=%v, got=%v\n", nums2, target2, expected2, res2)
}
