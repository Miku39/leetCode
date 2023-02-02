package main

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	got := canMakeArithmeticProgression([]int{3, 5, 1})
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
