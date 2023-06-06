package main

import "testing"

func TestMoveZeroes(t *testing.T) {
	got := moveZeroes([]int{0, 1, 0, 3, 12})
	want := []int{1, 3, 12, 0, 0}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMoveZeroes2(t *testing.T) {
	got := moveZeroes([]int{0})
	want := []int{0}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
