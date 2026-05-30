package main

func getBoxIndex(r, c int) int {
	ri := r / 3
	ci := c / 3
	return ri*3 + ci
}

func isValidSudoku(board [][]byte) bool {
	rowTrackingMap := make(map[int]map[rune]bool)
	colTrackingMap := make(map[int]map[rune]bool)
	boxTrackingMap := make(map[int]map[rune]bool)
	for r := range 9 {
		rowTrackingMap[r] = make(map[rune]bool)
	}
	for c := range 9 {
		colTrackingMap[c] = make(map[rune]bool)
	}

	for r := range 3 {
		for c := range 3 {
			index := r*3 + c
			boxTrackingMap[index] = make(map[rune]bool)
		}
	}

	for r := range 9 {
		for c := range 9 {
			b := rune(board[r][c])
			boxIndex := getBoxIndex(r, c)
			// fmt.Printf("b=%q r=%d, c=%d, boxIndex=%d\n", b, r, c, boxIndex)
			if b == '.' {
				continue
			}

			if colTrackingMap[c][b] || rowTrackingMap[r][b] || boxTrackingMap[boxIndex][b] {
				// fmt.Printf("c=%d r=%d boxIndex=%dcolTracking=%t rowTracking=%t boxTracking=%t\n",
				// 	c, r, boxIndex,
				// 	colTrackingMap[c][b],
				// 	rowTrackingMap[r][b],
				// 	boxTrackingMap[boxIndex][b],
				// )
				return false
			}
			colTrackingMap[c][b] = true
			rowTrackingMap[r][b] = true
			boxTrackingMap[boxIndex][b] = true
		}
	}

	return true
}
