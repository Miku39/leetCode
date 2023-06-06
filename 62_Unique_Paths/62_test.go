package main

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	got := uniquePaths(3, 7)
	want := 28

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
