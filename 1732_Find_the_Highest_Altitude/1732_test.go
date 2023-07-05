package main

import "testing"

func TestLargestAltitude(t *testing.T) {
	input := []int{-5, 1, 5, 0, -7}
	got := largestAltitude(input)
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
