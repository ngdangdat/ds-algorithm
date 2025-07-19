package main

import (
	"fmt"
)

func hasMatch(s string, dictString string) bool {
	idx := 0
	for _, c := range s {
		if idx < len(dictString) && c == rune(dictString[idx]) {
			idx += 1
		}

		if idx >= len(dictString) {
			return true
		}
	}

	return false
}

func findLongestWord(s string, dictionary []string) string {
	matchLength := 0
	result := ""
	sLength := len(s)
	for _, dictWord := range dictionary {
		wordLength := len(dictWord)
		if wordLength > sLength {
			continue
		}
		wordMatch := hasMatch(s, dictWord)
		if wordMatch && wordLength == matchLength && dictWord < result {
			result = dictWord
			continue
		}
		if wordMatch && wordLength > matchLength {
			result = dictWord
			matchLength = wordLength
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", ("abc" < "abe"))
	input1 := "abpcplea"
	dict1 := []string{"ale", "apple", "monkey", "plea"}
	res1 := findLongestWord(input1, dict1)
	fmt.Printf("Example 1, input: %v, dict: %v, result: %v\n", input1, dict1, res1)
}
