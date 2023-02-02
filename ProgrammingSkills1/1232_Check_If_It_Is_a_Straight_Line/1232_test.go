package main

import "testing"

func TestCheckStraightLine(t *testing.T) {

	input := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	got := checkStraightLine(input)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCheckStraightLine2(t *testing.T) {

	input := [][]int{{1, -8}, {2, -3}, {1, 2}}
	got := checkStraightLine(input)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCheckStraightLine3(t *testing.T) {

	input := [][]int{{4, 8}, {-2, 8}, {1, 8}, {8, 8}, {-5, 8}, {0, 8}, {7, 8}, {5, 8}}
	got := checkStraightLine(input)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
