package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	got := climbStairs(5)
	want := 8

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
