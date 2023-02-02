package main

import "testing"

func TestCountOdds(t *testing.T) {
	got := countOdds(3, 7)
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
