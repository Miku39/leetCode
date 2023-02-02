package main

import "testing"

func TestNearestValidPoint(t *testing.T) {
	points := [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}
	got := nearestValidPoint(3, 4, points)
	want := 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
