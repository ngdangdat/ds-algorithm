package main

func spiralOrder(matrix [][]int) []int {
	numRow := len(matrix)
	numCol := len(matrix[0])
	total := numRow * numCol
	res := make([]int, total)
	top, bottom, right, left := 0, numRow-1, numCol-1, 0
	index := 0
	for (left <= right) && (top <= bottom) {
		for c := left; c <= right; c++ {
			res[index] = matrix[top][c]
			index++
		}
		top++
		for r := top; r <= bottom; r++ {
			res[index] = matrix[r][right]
			index++
		}
		right--
		if top <= bottom {
			for c := right; c >= left; c-- {
				res[index] = matrix[bottom][c]
				index++
			}
			bottom--
		}
		if left <= right {
			for r := bottom; r >= top; r-- {
				res[index] = matrix[r][left]
				index++
			}
			left++
		}
	}
	return res
}
