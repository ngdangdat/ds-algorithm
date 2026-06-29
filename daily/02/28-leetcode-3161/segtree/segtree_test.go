package segtree

import (
	"math/rand"
	"testing"
)

// build creates a tree from an initial slice by Set-ing each element.
func build(vals []int) *MaxSegTree {
	t := NewMaxSegTree(len(vals))
	for i, v := range vals {
		t.Set(i, v)
	}
	return t
}

func TestMaxRangeBasic(t *testing.T) {
	// a = [2, 8, 1, 5]
	st := build([]int{2, 8, 1, 5})

	cases := []struct {
		l, r, want int
	}{
		{0, 0, 2},
		{1, 1, 8},
		{0, 1, 8},
		{2, 3, 5},
		{0, 3, 8},
		{2, 2, 1},
		{1, 3, 8},
	}
	for _, c := range cases {
		if got := st.MaxRange(c.l, c.r); got != c.want {
			t.Errorf("MaxRange(%d,%d) = %d, want %d", c.l, c.r, got, c.want)
		}
	}
}

func TestPrefixMax(t *testing.T) {
	// Prefix max from 0 is exactly how the real problem uses this tree.
	st := build([]int{2, 8, 1, 5})
	prefix := []int{2, 8, 8, 8} // MaxRange(0, i) for i = 0..3
	for i, want := range prefix {
		if got := st.MaxRange(0, i); got != want {
			t.Errorf("MaxRange(0,%d) = %d, want %d", i, got, want)
		}
	}
}

func TestSetUpdatesMax(t *testing.T) {
	st := build([]int{2, 8, 1, 5})

	// Lower the current max (index 1: 8 -> 0). New overall max should be 5.
	st.Set(1, 0)
	if got := st.MaxRange(0, 3); got != 5 {
		t.Errorf("after Set(1,0): MaxRange(0,3) = %d, want 5", got)
	}

	// Raise index 2 to a new maximum.
	st.Set(2, 9)
	if got := st.MaxRange(0, 3); got != 9 {
		t.Errorf("after Set(2,9): MaxRange(0,3) = %d, want 9", got)
	}
	if got := st.MaxRange(0, 1); got != 2 {
		t.Errorf("after updates: MaxRange(0,1) = %d, want 2", got)
	}
}

func TestEmptyRangeReturnsZero(t *testing.T) {
	st := build([]int{3, 1, 4})
	if got := st.MaxRange(2, 1); got != 0 {
		t.Errorf("empty range MaxRange(2,1) = %d, want 0", got)
	}
}

func TestAllZerosByDefault(t *testing.T) {
	st := NewMaxSegTree(5)
	if got := st.MaxRange(0, 4); got != 0 {
		t.Errorf("fresh tree MaxRange(0,4) = %d, want 0", got)
	}
}

func TestSingleElement(t *testing.T) {
	st := NewMaxSegTree(1)
	st.Set(0, 42)
	if got := st.MaxRange(0, 0); got != 42 {
		t.Errorf("MaxRange(0,0) = %d, want 42", got)
	}
}

// TestAgainstBruteForce fuzzes Set/MaxRange against a plain array oracle.
func TestAgainstBruteForce(t *testing.T) {
	const n = 64
	rng := rand.New(rand.NewSource(1))

	arr := make([]int, n)
	st := NewMaxSegTree(n)

	bruteMax := func(l, r int) int {
		if l > r {
			return 0
		}
		m := arr[l]
		for i := l + 1; i <= r; i++ {
			if arr[i] > m {
				m = arr[i]
			}
		}
		return m
	}

	for op := 0; op < 5000; op++ {
		if rng.Intn(2) == 0 {
			i := rng.Intn(n)
			v := rng.Intn(1000)
			arr[i] = v
			st.Set(i, v)
		} else {
			l := rng.Intn(n)
			r := rng.Intn(n)
			if l > r {
				l, r = r, l
			}
			want := bruteMax(l, r)
			if got := st.MaxRange(l, r); got != want {
				t.Fatalf("op %d: MaxRange(%d,%d) = %d, want %d", op, l, r, got, want)
			}
		}
	}
}
