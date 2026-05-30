package main

func setZeroes(matrix [][]int) {
	maxR := len(matrix)
	maxC := len(matrix[0])
	firstRowHasZero := false
	firstColHasZero := false
	for r, cs := range matrix {
		for c, v := range cs {
			if v == 0 {
				if r == 0 {
					firstRowHasZero = true
				}
				if c == 0 {
					firstColHasZero = true
				}
				matrix[0][c] = 0
				matrix[r][0] = 0
			}
		}
	}
	for r := 1; r < maxR; r++ {
		if matrix[r][0] == 0 {
			for c := range maxC {
				matrix[r][c] = 0
			}
		}
	}
	for c := 1; c < maxC; c++ {
		if matrix[0][c] == 0 {
			for r := range maxR {
				matrix[r][c] = 0
			}
		}
	}
	if firstRowHasZero {
		for c := range maxC {
			matrix[0][c] = 0
		}
	}
	if firstColHasZero {
		for r := range maxR {
			matrix[r][0] = 0
		}

	}
}
