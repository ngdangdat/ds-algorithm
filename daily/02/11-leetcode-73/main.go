package main

func setRowColZeroes(matrix [][]int, visited map[[2]int]bool, r, c, maxR, maxC int) {
	for er := range maxR {
		if visited[[2]int{er, c}] {
			continue
		}
		if matrix[er][c] != 0 {
			visited[[2]int{er, c}] = true
		}
		matrix[er][c] = 0
	}
	for ec := range maxC {
		if visited[[2]int{r, ec}] {
			continue
		}
		if matrix[r][ec] != 0 {
			visited[[2]int{r, ec}] = true
		}
		matrix[r][ec] = 0
	}
}

func setZeroes(matrix [][]int) {
	visited := make(map[[2]int]bool)
	maxR := len(matrix)
	maxC := len(matrix[0])
	for r, cs := range matrix {
		for c, v := range cs {
			if visited[[2]int{r, c}] {
				continue
			}
			if v == 0 {
				setRowColZeroes(matrix, visited, r, c, maxR, maxC)
			}
		}
	}
}
