package main

import "testing"

func TestArraySign(t *testing.T) {
	input := []int{-1, -2, -3, -4, 3, 2, 1}
	got := arraySign(input)
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
