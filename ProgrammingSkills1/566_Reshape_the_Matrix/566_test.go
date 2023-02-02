package main

import "testing"

func TestMatrixReshape(t *testing.T) {
	got := matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4)
	want := [][]int{{1, 2, 3, 4}}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMatrixReshape2(t *testing.T) {
	got := matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4)
	want := [][]int{{1, 2}, {3, 4}}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b [][]int) bool {
	for i, v := range a {
		for j, w := range v {
			if w != b[i][j] {
				return false
			}
		}
	}
	return true
}
