package main

import "testing"

func TestPivotIndex(t *testing.T) {
	got := pivotIndex([]int{1, 7, 3, 6, 5, 6})
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
