package main

// maxArea returns the maximum amount of water a container can store, where
// height[i] is the height of the i-th vertical line. The container is formed
// by two lines and the x-axis; its area is min(height[i], height[j]) * (j - i).
//
// TODO: implement.
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxWater := 0
	for l < r {
		hL, hR := height[l], height[r]
		water := min(hL, hR) * (r - l)
		if water > maxWater {
			maxWater = water
		}
		if hL < hR {
			l++
		} else {
			r--
		}
	}
	return maxWater
}
