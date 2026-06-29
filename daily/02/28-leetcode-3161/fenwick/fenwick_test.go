package fenwick

import (
	"math/rand"
	"sort"
	"testing"
)

func build(n int, present ...int) *CountFenwick {
	f := NewCountFenwick(n)
	for _, i := range present {
		f.Add(i)
	}
	return f
}

func TestPrefixCount(t *testing.T) {
	// present at {1, 3, 4} over indices 0..5
	f := build(6, 1, 3, 4)
	cases := []struct{ i, want int }{
		{-1, 0},
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 3},
	}
	for _, c := range cases {
		if got := f.PrefixCount(c.i); got != c.want {
			t.Errorf("PrefixCount(%d) = %d, want %d", c.i, got, c.want)
		}
	}
}

func TestFloor(t *testing.T) {
	f := build(8, 2, 5) // present at 2 and 5
	cases := []struct{ i, want int }{
		{0, -1}, // nothing <= 0
		{1, -1},
		{2, 2}, // exactly on a present index
		{3, 2},
		{4, 2},
		{5, 5},
		{7, 5},
	}
	for _, c := range cases {
		if got := f.Floor(c.i); got != c.want {
			t.Errorf("Floor(%d) = %d, want %d", c.i, got, c.want)
		}
	}
}

func TestPredecessor(t *testing.T) {
	f := build(8, 2, 5)
	cases := []struct{ i, want int }{
		{0, -1},
		{2, -1}, // strictly less than 2 -> none
		{3, 2},
		{5, 2}, // strictly less than 5 -> 2
		{6, 5},
		{7, 5},
	}
	for _, c := range cases {
		if got := f.Predecessor(c.i); got != c.want {
			t.Errorf("Predecessor(%d) = %d, want %d", c.i, got, c.want)
		}
	}
}

func TestSuccessor(t *testing.T) {
	f := build(8, 2, 5)
	cases := []struct{ i, want int }{
		{0, 2},
		{1, 2},
		{2, 5}, // strictly greater than 2 -> 5
		{4, 5},
		{5, -1}, // strictly greater than 5 -> none
		{7, -1},
	}
	for _, c := range cases {
		if got := f.Successor(c.i); got != c.want {
			t.Errorf("Successor(%d) = %d, want %d", c.i, got, c.want)
		}
	}
}

func TestEmpty(t *testing.T) {
	f := NewCountFenwick(5)
	if f.PrefixCount(4) != 0 {
		t.Errorf("empty PrefixCount = %d, want 0", f.PrefixCount(4))
	}
	if f.Floor(3) != -1 || f.Predecessor(3) != -1 || f.Successor(3) != -1 {
		t.Errorf("empty lookups should all be -1")
	}
}

func TestSingle(t *testing.T) {
	f := build(4, 2)
	if f.Floor(2) != 2 || f.Floor(3) != 2 || f.Floor(1) != -1 {
		t.Errorf("Floor wrong for single element")
	}
	if f.Predecessor(2) != -1 || f.Successor(2) != -1 {
		t.Errorf("single element has no pred/succ relative to itself")
	}
	if f.Successor(1) != 2 || f.Predecessor(3) != 2 {
		t.Errorf("pred/succ around single element wrong")
	}
}

// TestAgainstBruteForce fuzzes against a plain sorted-set oracle.
func TestAgainstBruteForce(t *testing.T) {
	const n = 50
	rng := rand.New(rand.NewSource(7))
	f := NewCountFenwick(n)
	present := map[int]bool{}

	sortedKeys := func() []int {
		ks := make([]int, 0, len(present))
		for k := range present {
			ks = append(ks, k)
		}
		sort.Ints(ks)
		return ks
	}
	bruteFloor := func(i int) int {
		ks := sortedKeys()
		ans := -1
		for _, k := range ks {
			if k <= i {
				ans = k
			}
		}
		return ans
	}
	brutePred := func(i int) int {
		ks := sortedKeys()
		ans := -1
		for _, k := range ks {
			if k < i {
				ans = k
			}
		}
		return ans
	}
	bruteSucc := func(i int) int {
		ks := sortedKeys()
		for _, k := range ks {
			if k > i {
				return k
			}
		}
		return -1
	}
	bruteCount := func(i int) int {
		c := 0
		for k := range present {
			if k <= i {
				c++
			}
		}
		return c
	}

	for op := 0; op < 4000; op++ {
		switch rng.Intn(5) {
		case 0:
			i := rng.Intn(n)
			if !present[i] {
				present[i] = true
				f.Add(i)
			}
		case 1:
			i := rng.Intn(n)
			if got, want := f.PrefixCount(i), bruteCount(i); got != want {
				t.Fatalf("op %d: PrefixCount(%d)=%d want %d", op, i, got, want)
			}
		case 2:
			i := rng.Intn(n)
			if got, want := f.Floor(i), bruteFloor(i); got != want {
				t.Fatalf("op %d: Floor(%d)=%d want %d", op, i, got, want)
			}
		case 3:
			i := rng.Intn(n)
			if got, want := f.Predecessor(i), brutePred(i); got != want {
				t.Fatalf("op %d: Predecessor(%d)=%d want %d", op, i, got, want)
			}
		case 4:
			i := rng.Intn(n)
			if got, want := f.Successor(i), bruteSucc(i); got != want {
				t.Fatalf("op %d: Successor(%d)=%d want %d", op, i, got, want)
			}
		}
	}
}
