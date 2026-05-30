package main

import (
	"slices"
)

func classifyStr(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	groupMap := make(map[string][]string)
	for _, s := range strs {
		g := classifyStr(s)
		groupMap[g] = append(groupMap[g], s)
	}
	for _, gs := range groupMap {
		each := []string{}
		for _, eg := range gs {
			each = append(each, eg)
		}
		res = append(res, each)
	}
	return res
}

// func main() {
// 	str := "dat"
// 	strs := []rune(str)
// 	slices.Sort(strs)
// 	fmt.Printf("str=%v", string(strs))
// }
