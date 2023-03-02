package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTwoSum2(t *testing.T) {
	got := twoSum([]int{3, 2, 4}, 6)
	want := []int{1, 2}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
