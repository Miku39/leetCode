package main

import "testing"

func TestSearch(t *testing.T) {
	got := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSearch2(t *testing.T) {
	got := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	want := -1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
