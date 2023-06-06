package main

import "testing"

func TestIsHappy(t *testing.T) {
	got := isHappy(19)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
