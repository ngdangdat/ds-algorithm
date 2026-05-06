package main

import "slices"

func findWordsContaining(words []string, x byte) []int {
	out := make([]int, 0, len(words))
	for i, w := range words {
		if slices.Contains([]byte(w), x) {
			out = append(out, i)
		}
	}
	return out
}

