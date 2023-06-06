package main

import "testing"

func TestFloodFill(t *testing.T) {
	got := floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2)
	want := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range v {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
