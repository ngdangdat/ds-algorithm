package main

import (
	"fmt"
)

func getHoursToEat(piles []int, speed int) int {
	res := 0
	for _, pile := range piles {
		res += (pile + speed - 1) / speed
		// fmt.Printf("pile=%d, speed=%d, res=%d\n", pile, speed, res)
	}
	return res
}

func subMinEatingSpeed(piles []int, currSpeed, targetHour, l, r int) int {
	// fmt.Printf("enter, l=%d, r=%d\n", l, r)
	if l > r {
		return currSpeed
	}
	if l == r {
		hours := getHoursToEat(piles, l)
		if hours <= targetHour && (l < currSpeed || currSpeed < 0) {
			return l
		}
		return currSpeed
	}
	mid := l + (r-l)/2
	hours := getHoursToEat(piles, mid)
	if hours > targetHour {
		// need to be faster
		return subMinEatingSpeed(piles, currSpeed, targetHour, mid+1, r)
	}
	// could be slower
	if currSpeed < 0 || currSpeed > mid {
		currSpeed = mid
	}
	return subMinEatingSpeed(piles, currSpeed, targetHour, l, mid-1)
}

func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, pile := range piles {
		if pile >= maxPile {
			maxPile = pile
		}
	}
	return subMinEatingSpeed(piles, -1, h, 1, maxPile)
}

func main() {
	piles1 := []int{3, 6, 7, 11}
	hours1 := 8
	expected1 := 4

	result1 := minEatingSpeed(piles1, hours1)
	fmt.Printf("ex1 piles=%v, hours=%v, expected=%v, got=%v\n", piles1, hours1, expected1, result1)

	piles2 := []int{30, 11, 23, 4, 20}
	hours2 := 5
	expected2 := 30

	result2 := minEatingSpeed(piles2, hours2)
	fmt.Printf("ex1 piles=%v, hours=%v, expected=%v, got=%v\n", piles2, hours2, expected2, result2)

	piles3 := []int{312884470}
	hours3 := 312884469
	expected3 := 2

	result3 := minEatingSpeed(piles3, hours3)
	fmt.Printf("ex1 piles=%v, hours=%v, expected=%v, got=%v\n", piles3, hours3, expected3, result3)
}
