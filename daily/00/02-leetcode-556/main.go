package main

import (
	"fmt"
	"math"
	"slices"
)

func getDigits(n int) []int {
	digits := []int{}
	num := n
	for num != 0 {
		d := num % 10
		num /= 10
		digits = append(digits, d)
	}
	slices.Reverse(digits)
	return digits
}

func getIntFromDigits(digits []int) int {
	exp := 0
	idx := len(digits) - 1
	num := 0
	for idx >= 0 {
		mul := (int)(math.Pow10(exp))
		
		num += mul * digits[idx]
		exp += 1
		idx -= 1
	}
	return num
}

func nextGreaterElement(n int) int {
	digits := getDigits(n)
	idx := len(digits) - 1
	pivot := idx
	for idx > 0 {
		if digits[idx-1] < digits[idx] {
			pivot = idx - 1
			break
		}
		idx--
	}

	for idx = len(digits) - 1; idx > pivot; {
		if digits[idx] > digits[pivot] {
			digits[idx], digits[pivot] = digits[pivot], digits[idx]
			break
		}
		idx--
	}
	slices.Reverse(digits[pivot+1:])

	newNum := getIntFromDigits(digits)
	fmt.Printf("newNum: %d\n", newNum)
	if newNum <= math.MaxInt32 && newNum > n {
		return newNum
	}

	return -1
}

func main() {
	n := 21
	expected := -1
	nextEl := nextGreaterElement(n)
	fmt.Printf("n: %d, nextEl: %d, correct: %d\n", n, nextEl, expected)
	// fmt.Printf("maxInt: %d, maxInt32: %d\n", math.MaxInt, math.MaxInt32)
}
