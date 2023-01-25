package main

import "testing"

func TestMaximumWealth(t *testing.T) {
	got := maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}})
	want := 6

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMaximumWealth2(t *testing.T) {
	got := maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}})
	want := 10

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMaximumWealth3(t *testing.T) {
	got := maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}})
	want := 17

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
