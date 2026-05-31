package main

func getNextVisit(numRow, numCol, r, c int) (int, int, bool) {
	if r == 0 {
		if c == numCol-1 {
			return r + 1, c, false
		}
		return r, c + 1, false
	}
	if c == numCol-1 {
		if r == numRow-1 {
			return r, c - 1, false
		}
		return r + 1, c, false
	}
	if r == numRow-1 {
		if c == 0 {
			return r - 1, c, false
		}
		return r, c - 1, false
	}
	if r == 1 {
		return 0, 0, true
	}
	return r - 1, c, false
}

func spiralOrder(matrix [][]int) []int {
	visitedCount := 0
	numRow := len(matrix)
	numCol := len(matrix[0])
	total := numRow * numCol
	buffer, r, c := 0, 0, 0
	res := make([]int, total)
	isEnd := false
	for visitedCount < total {
		res[visitedCount] = matrix[r][c]
		r, c, isEnd = getNextVisit(numRow, numCol, r-buffer, c-buffer)
		if isEnd {
			numRow -= 2
			numCol -= 2
			buffer += 1
		}
		r += buffer
		c += buffer
		visitedCount += 1
	}
	return res
}
