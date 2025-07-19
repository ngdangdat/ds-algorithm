package main

import "testing"

func TestLongestWordInDict(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		dictionary []string
		expected   string
	}{
		{
			name:  "normal case 1",
			input: "abpcplea",
			dictionary: []string{
				"ale",
				"apple",
				"monkey",
				"plea",
			},
			expected: "apple",
		},
		{
			name:  "normal case 2",
			input: "abce",
			dictionary: []string{
				"abe",
				"abc",
			},
			expected: "abc",
		},
		{
			name:       "single character words",
			input:      "abpcplea",
			dictionary: []string{"a", "b", "c"},
			expected:   "a",
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			got := findLongestWord(tCase.input, tCase.dictionary)
			if got != tCase.expected {
				t.Errorf("findLongestWord expected: %v, got: %v", tCase.expected, got)
			}
		})
	}
}
