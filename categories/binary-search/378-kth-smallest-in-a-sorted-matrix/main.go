package main

import "fmt"

func insertionSort(mainSlice []int, elSlice []int) {
	index := 0
	for _, el := range elSlice {
		insertIndex := -1
		for index < len(mainSlice) {
			curr := mainSlice[index]
			if curr > el {
				insertIndex = index
				break
			}
			index += 1
		}

	}
}

func kthSmallest(matrix [][]int, k int) int {
	flattenSlice := []int{}

	return -1
}

func main() {
	matrix1 := [][]int{
		{1, 5, 9},
		{1, 11, 13},
		{2, 13, 15},
	}
	k1 := 3
	expected1 := 2
	got1 := kthSmallest(matrix1, k1)

	fmt.Printf("Case 1; matrix=%v, k=%v, expected=%v, got=%v\n", matrix1, k1, expected1, got1)
}
