package main

func isAnagram(s string, t string) bool {
	countMap := make(map[rune]int)
	for _, es := range s {
		countMap[es] += 1
	}
	for _, et := range t {
		ct, ok := countMap[et]
		if !ok {
			return false
		}
		countMap[et] = ct - 1
		if countMap[et] < 0 {
			return false
		}
	}
	for _, v := range countMap {
		if v > 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	sCountMap := make(map[rune]int)
	tCountMap := make(map[rune]int)

	for _, es := range s {
		sCountMap[es] += 1
	}

	for _, et := range t {
		tCountMap[et] += 1
		sc, ok := sCountMap[et]
		if !ok {
			return false
		}
		if tCountMap[et] > sc {
			return false
		}
	}
	for k, cs := range sCountMap {
		ct, ok := tCountMap[k]
		if !ok {
			return false
		}
		if ct != cs {
			return false
		}
	}
	return true
}
