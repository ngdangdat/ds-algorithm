package main

import "fmt"

func binarySearch(nums []int, target int, l int, r int) bool {
	if l > r {
		return false
	}

	mid := l + (r-l)/2
	if nums[mid] < target {
		return binarySearch(nums, target, mid+1, r)
	} else if nums[mid] > target {
		return binarySearch(nums, target, l, mid-1)
	}

	return true
}

func binarySearchMatrix(matrix [][]int, rowLen, target, l, r int) bool {
	if l > r {
		return false
	}
	mid := l + (r-l)/2
	if target < matrix[mid][0] {
		return binarySearchMatrix(matrix, rowLen, target, l, mid-1)
	} else if target > matrix[mid][rowLen-1] {
		return binarySearchMatrix(matrix, rowLen, target, mid+1, r)
	}

	return binarySearch(matrix[mid], target, 0, rowLen-1)
}

func searchMatrix(matrix [][]int, target int) bool {
	rowLen := len(matrix[0])
	return binarySearchMatrix(matrix, rowLen, target, 0, len(matrix)-1)
}

func main() {
	matrix1 := [][]int{{1, 4, 7, 11}, {2, 5, 8, 12}, {3, 6, 9, 16}, {10, 13, 14, 17}}
	target1 := 5
	expected1 := true
	got1 := searchMatrix(matrix1, target1)
	fmt.Printf("ex1, matrix=%v, target=%v, expected=%v, got=%v\n", matrix1, target1, expected1, got1)
}
