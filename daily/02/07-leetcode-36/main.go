package main

func getBoxIndex(r, c int) int {
	ri := r / 3
	ci := c / 3
	return ri*3 + ci
}

func isValidSudoku(board [][]byte) bool {
	rowTrackingMap := [9][9]bool{}
	colTrackingMap := [9][9]bool{}
	boxTrackingMap := [9][9]bool{}

	for r := range 9 {
		for c := range 9 {
			if rune(board[r][c]) == '.' {
				continue
			}
			b := int(board[r][c]-'0') - 1
			boxIndex := getBoxIndex(r, c)

			if colTrackingMap[c][b] || rowTrackingMap[r][b] || boxTrackingMap[boxIndex][b] {
				return false
			}
			colTrackingMap[c][b] = true
			rowTrackingMap[r][b] = true
			boxTrackingMap[boxIndex][b] = true
		}
	}

	return true
}
