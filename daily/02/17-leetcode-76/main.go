package main

import (
	"math"
)

func minWindow(s string, t string) string {
	fitMap := make(map[byte]int)
	matchMap := make(map[byte]int)
	for _, cT := range t {
		fitMap[byte(cT)] += 1
	}
	required := len(fitMap)
	left, right := 0, 0
	bestStart, bestLength := 0, math.MaxInt
	formed := 0
	for right < len(s) {
		cS := s[right]
		count, ok := fitMap[byte(cS)]
		if ok {
			matchMap[byte(cS)]++
			if matchMap[byte(cS)] == count {
				formed += 1
			}
		}
		for formed == required {
			currLength := right - left + 1
			if currLength < bestLength {
				bestLength = currLength
				bestStart = left
			}
			if _, ok := matchMap[byte(s[left])]; ok {
				matchMap[byte(s[left])]--
				if matchMap[s[left]] < fitMap[s[left]] {
					formed--
				}
			}
			left++
		}
		right++
	}
	if bestLength >= len(s)+1 {
		return ""
	}
	return s[bestStart : bestStart+bestLength]
}
