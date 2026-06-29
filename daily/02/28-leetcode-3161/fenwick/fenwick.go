package fenwick

// CountFenwick is a Fenwick (binary indexed) tree over indices 0..n-1, where
// each index is either absent (0) or present (1). It supports:
//
//   - Add(i):           mark index i as present                 -- O(log n)
//   - PrefixCount(i):   how many present indices are in [0, i]   -- O(log n)
//   - Predecessor(i):   largest present index j with j <  i      -- O(log n), -1 if none
//   - Successor(i):     smallest present index j with j >  i     -- O(log n), -1 if none
//   - Floor(i):         largest present index j with j <= i      -- O(log n), -1 if none
//
// Indices are 0-based on the public API. (Internally a Fenwick tree is usually
// 1-indexed; how you bridge that is up to you.)
//
// This is a learning sandbox: implement the methods so fenwick_test.go passes.
// In the real problem this replaces the O(n) sorted `blocks` slice:
//   - L (left neighbor)  = Predecessor(idx)
//   - R (right neighbor) = Successor(idx)
//   - floor(x)           = Floor(indexMap[x])
type CountFenwick struct {
	// TODO: choose your representation.
	//
	// Classic layout: a 1-indexed []int `bit` of length n+1. A point update at
	// 1-indexed position p walks p += p & (-p); a prefix query walks p -= p & (-p).
	// Predecessor/Successor/Floor can be derived from PrefixCount (binary search)
	// or via a binary-lifting walk on the tree for a tighter O(log n).
	n int
}

// NewCountFenwick returns a tree over indices 0..n-1, all absent.
func NewCountFenwick(n int) *CountFenwick {
	// TODO: allocate storage, store n.
	return &CountFenwick{n: n}
}

// Add marks index i (0-based) as present.
func (f *CountFenwick) Add(i int) {
	// TODO: implement.
}

// PrefixCount returns how many present indices lie in [0, i] (inclusive).
// If i < 0, return 0.
func (f *CountFenwick) PrefixCount(i int) int {
	// TODO: implement.
	return 0
}

// Predecessor returns the largest present index strictly less than i, or -1.
func (f *CountFenwick) Predecessor(i int) int {
	// TODO: implement.
	return -1
}

// Successor returns the smallest present index strictly greater than i, or -1.
func (f *CountFenwick) Successor(i int) int {
	// TODO: implement.
	return -1
}

// Floor returns the largest present index less than or equal to i, or -1.
func (f *CountFenwick) Floor(i int) int {
	// TODO: implement.
	return -1
}

func main() {
	fenwickTree := NewCountFenwick(10)

}
