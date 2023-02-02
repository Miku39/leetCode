package main

import "testing"

func TestRunningSum(t *testing.T) {
	got := runningSum([]int{1, 2, 3, 4})
	want := []int{1, 3, 6, 10}

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
