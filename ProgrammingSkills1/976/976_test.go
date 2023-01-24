package main

import "testing"

func TestLargestPerimeter(t *testing.T) {
	got := largestPerimeter([]int{2, 1, 2})
	want := 5

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
