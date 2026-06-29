package main

import (
	"reflect"
	"testing"
)

func TestGetResults(t *testing.T) {
	tests := []struct {
		name    string
		queries [][]int
		want    []bool
	}{
		{
			name:    "example1",
			queries: [][]int{{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2}},
			want:    []bool{false, true, true},
		},
		{
			name:    "example2",
			queries: [][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}},
			want:    []bool{true, true, false},
		},
		{
			name:    "no_obstacles_full_range",
			queries: [][]int{{2, 5, 5}, {2, 5, 1}},
			want:    []bool{true, true},
		},
		{
			name:    "block_exactly_fits_touching",
			queries: [][]int{{1, 4}, {2, 4, 4}},
			want:    []bool{true},
		},
		{
			name:    "gap_after_obstacle",
			queries: [][]int{{1, 3}, {2, 10, 7}, {2, 10, 8}},
			want:    []bool{true, false},
		},
		{
			name:    "no_type2_queries",
			queries: [][]int{{1, 1}, {1, 5}},
			want:    []bool{},
		},
		{
			name:    "obstacle_at_query_bound",
			queries: [][]int{{1, 5}, {2, 5, 5}, {1, 2}, {2, 5, 4}},
			want:    []bool{true, false},
		},
		{
			// Guards against slice-aliasing: a type-2 query must not mutate the
			// obstacle list. The first query's boundary append must not poison
			// the obstacles seen by the second.
			name:    "type2_must_not_mutate_obstacles",
			queries: [][]int{{1, 12}, {1, 1}, {1, 14}, {2, 6, 1}, {2, 9, 6}},
			want:    []bool{true, true},
		},
		{
			// A large gap living to the RIGHT of x must not leak into the answer.
			// Obstacles at 2 and 20 -> gap[20]=18 is far past x=5. Within [0,5]
			// the gaps are [0,2]=2 and trailing [2,5]=3, so max gap 3 < 4 -> false.
			// Catches a range-max query that ignores its [0, x] bound.
			name:    "gap_beyond_x_excluded",
			queries: [][]int{{1, 2}, {1, 20}, {2, 5, 4}},
			want:    []bool{false},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getResults(tc.queries)
			if len(tc.want) == 0 && len(got) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("getResults(%v) = %v, want %v", tc.queries, got, tc.want)
			}
		})
	}
}

// TestGetResultsPreservesInput guards that the input queries are not mutated.
func TestGetResultsPreservesInput(t *testing.T) {
	queries := [][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}}
	snapshot := make([][]int, len(queries))
	for i, q := range queries {
		snapshot[i] = append([]int(nil), q...)
	}

	getResults(queries)

	for i := range queries {
		if !reflect.DeepEqual(queries[i], snapshot[i]) {
			t.Errorf("input mutated at query %d: got %v, want %v", i, queries[i], snapshot[i])
		}
	}
}
