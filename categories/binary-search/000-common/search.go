package main

import "fmt"

func biSearchRange(ns []int, n, l, r int) int {
	if l > r {
		return -1
	}

	mid := l + (r-l)/2
	if ns[mid] > n {
		return biSearchRange(ns, n, l, mid-1)
	} else if ns[mid] < n {
		return biSearchRange(ns, n, mid+1, r)
	}
	return mid
}

func biSearch(ns []int, n int) int {
	return biSearchRange(ns, n, 0, len(ns)-1)
}

func main() {
	ns1 := []int{1, 2, 3, 4, 9, 10, 25}
	n1 := 4

	res1 := biSearch(ns1, n1)
	fmt.Printf("ex1, ns=%v, n=%v, res=%v\n", ns1, n1, res1)
}
