package main

import (
	"sort"
)

type MaxGapSegTree struct {
	// underlying array of blocks
	nodes []int
	n     int
}

func NewMaxGapSegtree(n int) MaxGapSegTree {
	l := 2
	for l <= n {
		l *= 2
	}
	nodes := make([]int, 2*l)
	return MaxGapSegTree{n: l, nodes: nodes}
}

func insertPos(blocks []int, block int) int {
	return sort.SearchInts(blocks, block)
}
func insertToBlocks(blocks []int, pos, block int) []int {
	blocks = append(blocks, 0)
	copy(blocks[pos+1:], blocks[pos:])
	blocks[pos] = block
	return blocks
}

func (t *MaxGapSegTree) Update(index, value int) {
	pos := t.n + index
	t.nodes[pos] = value
	for pos > 1 {
		pos = pos / 2
		t.nodes[pos] = max(t.nodes[pos*2], t.nodes[pos*2+1])
	}
}

func (t *MaxGapSegTree) MaxGapSize(l, r int) int {
	res := 0
	l += t.n
	r += t.n + 1
	for l < r {
		if l&1 == 1 {
			res = max(res, t.nodes[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t.nodes[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func getResults(queries [][]int) []bool {
	res := []bool{}
	elSet := map[int]bool{}
	for _, q := range queries {
		el := q[1]
		elSet[el] = true
	}
	nums := make([]int, 0, len(elSet))
	for k := range elSet {
		nums = append(nums, k)
	}
	sort.Ints(nums)
	indexMap := map[int]int{}
	for i, n := range nums {
		indexMap[n] = i
	}
	blocks := []int{0} // 0 is always there
	segTree := NewMaxGapSegtree(len(nums))

	for _, q := range queries {
		qType := q[0]
		if qType == 1 {
			//update
			val := q[1]
			// we need to update the right part here, too
			pos := insertPos(blocks, val)
			blocks = insertToBlocks(blocks, pos, val)
			prev := blocks[pos-1]
			gap := val - prev
			segTree.Update(indexMap[val], gap)
			if pos < len(blocks)-1 {
				segTree.Update(indexMap[blocks[pos+1]], blocks[pos+1]-val)
			}
		} else {
			//query
			q, sz := q[1], q[2]
			rightMostPos := sort.SearchInts(blocks, q)
			trailingGap := q - blocks[rightMostPos-1]
			maxGap := max(segTree.MaxGapSize(0, indexMap[q]), trailingGap)
			possible := maxGap >= sz
			res = append(res, possible)
		}
	}

	return res
}
