package segtree

// MaxSegTree is a segment tree over a fixed-size integer array that supports:
//
//   - Set(i, v):        assign a[i] = v               -- O(log n)
//   - MaxRange(l, r):   max of a[l..r] inclusive       -- O(log n)
//
// All elements start at 0. Indices are 0-based. MaxRange uses an inclusive
// range [l, r].
//
// This is a learning sandbox: implement the three methods below so that
// segtree_test.go passes. Do NOT look at the LeetCode solution while doing it.
type MaxSegTree struct {
	// TODO: choose your representation.
	nodes []int
	//
	// A common layout: a flat []int of length 2*size, 1-indexed, where node i
	// has children 2i and 2i+1, the root is index 1, and the leaves occupy the
	// second half. `size` is the number of leaves (round n up to a power of two,
	// or just use n directly with the recursive style).
	n int
}

// NewMaxSegTree returns a tree backed by n elements, all initialized to 0.
func NewMaxSegTree(n int) *MaxSegTree {
	pow := 2
	for pow < n {
		pow *= 2
	}
	n = pow
	nodes := make([]int, 2*n+1)
	return &MaxSegTree{n: n, nodes: nodes}
}

// Set assigns a[i] = v and repairs the maxes on the path from the leaf to the
// root.
func (t *MaxSegTree) Set(i, v int) {
	pos := t.n + i
	t.nodes[pos] = v
	for pos > 1 {
		pos = pos / 2
		t.nodes[pos] = max(t.nodes[2*pos], t.nodes[2*pos+1])
	}
}

func (t *MaxSegTree) query(node, low, high, l, r int) int {
	if l > high || r < low {
		return 0
	}
	if low >= l && high <= r {
		return t.nodes[node]
	}
	mid := (high + low) / 2
	leftChild, rightChild := node*2, node*2+1
	return max(
		t.query(leftChild, low, mid, l, r),
		t.query(rightChild, mid+1, high, l, r),
	)
}

// MaxRange returns the maximum of a[l..r] (inclusive). If the range is empty
// (l > r), return 0.
func (t *MaxSegTree) MaxRange(l, r int) int {
	return t.query(1, 0, t.n-1, l, r)
}
