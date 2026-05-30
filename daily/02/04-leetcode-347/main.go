package main

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n] += 1
	}
	reverseCountMap := make(map[int][]int)
	fs := []int{}
	for n, frequency := range countMap {
		cm, ok := reverseCountMap[frequency]
		if !ok {
			fs = append(fs, frequency)
		}
		reverseCountMap[frequency] = append(cm, n)
	}
	taken := 0
	res := []int{}
	for i := len(nums); taken < k; i-- {
		cm, ok := reverseCountMap[i]
		if !ok {
			continue
		}
		for _, v := range cm {
			res = append(res, v)
			taken += 1
			if taken >= k {
				break
			}
		}
	}

	return res
}
