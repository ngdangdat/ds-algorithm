package main

func longestConsecutive(nums []int) int {
	checkMap := make(map[int]bool)
	maxCons := 0
	for _, n := range nums {
		checkMap[n] = true
	}
	for n := range checkMap {
		if !checkMap[n-1] {
			cons := 0
			j := n
			for checkMap[j] {
				j += 1
				cons += 1
			}
			if cons > maxCons {
				maxCons = cons
			}
		}
	}
	return maxCons
}
