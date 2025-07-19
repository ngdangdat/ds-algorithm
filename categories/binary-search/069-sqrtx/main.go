package main

import "fmt"

func biSectSqrtx(limit, l, r, curr int) int {
	if l > r {
		return curr
	}

	mid := l + (r-l)/2

	if mid*mid < limit {
		curr = mid
		return biSectSqrtx(limit, mid+1, r, curr)
	} else if mid*mid > limit {
		return biSectSqrtx(limit, l, mid-1, curr)
	}

	return mid
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	return biSectSqrtx(x, 1, x, -1)
}

func main() {
	n1 := 2
	fmt.Printf("ex1, n=%v, res=%v\n", n1, mySqrt(n1))

	n2 := 9
	fmt.Printf("ex2, n=%v, res=%v\n", n2, mySqrt(n2))
}
