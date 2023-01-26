package main

import "testing"

func TestDiagonalSum(t *testing.T) {
	got := diagonalSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	want := 25

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
