package main

import (
	"fmt"
)

type tuple = [2]int

func countCoveredBuildings(n int, buildings [][]int) int {
	minCorr, maxCorr := 1, n

	var xMap, yMap map[int]tuple = map[int]tuple{}, map[int]tuple{}
	res := 0
	for _, building := range buildings {
		x, y := building[0], building[1]
		if _, ok := xMap[x]; !ok {
			xMap[x] = tuple{maxCorr, minCorr}
		}
		if _, ok := yMap[y]; !ok {
			yMap[y] = tuple{maxCorr, minCorr}
		}
		if (y < xMap[x][0]) {
			xMap[x] = tuple{y, xMap[x][1]}
		}
		if (y > xMap[x][1]) {
			xMap[x] = tuple{xMap[x][0], y}
		}
		if (x < yMap[y][0]) {
			yMap[y] = tuple{x, yMap[y][1]}
		}
		if (x > yMap[y][1]) {
			yMap[y] = tuple{yMap[y][0], x}
		}
	}
	for _, building := range buildings {
		isXCovered := false
		isYCovered := false
		x := building[0]
		y := building[1]
		checkList := xMap[x]
		if y > checkList[0] && y < checkList[1] {
			isXCovered = true
		}
		checkList = yMap[y]
		if x > checkList[0] && x < checkList[1] {
			isYCovered = true
		}
		if isXCovered && isYCovered {
			res += 1
		}
	}
	// fmt.Printf("xMap: %v\n", xMap)
	// fmt.Printf("yMap: %v\n", yMap)
	return res
}

func main() {
	res := countCoveredBuildings(
		5,
		// [][]int{{1,1},{1,2},{2,1},{2,2}}, // 0
		// [][]int{{1,2},{2,2},{3,2},{2,1},{2,3}}, // 1
		[][]int{{1,3},{3,2},{3,3},{3,5},{5,3}}, // 1
	)
	fmt.Printf("result: %d\n", res)
}
