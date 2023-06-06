package main

import "testing"

func TestFirstBadVersion(t *testing.T) {
	got := firstBadVersion(5)
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
