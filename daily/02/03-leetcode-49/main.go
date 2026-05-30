package main

func classifyStr(s string) [26]int {
	res := [26]int{}
	for _, r := range []byte(s) {
		subR := r - byte('a')
		res[subR] += 1
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	groupMap := make(map[[26]int][]string)
	for _, s := range strs {
		g := classifyStr(s)
		groupMap[g] = append(groupMap[g], s)
	}
	for _, gs := range groupMap {
		res = append(res, gs)
	}
	return res
}

// func main() {
// 	str := "dat"
// 	strs := []rune(str)
// 	slices.Sort(strs)
// 	fmt.Printf("str=%v", string(strs))
// }
