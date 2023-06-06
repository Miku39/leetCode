package main

import (
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	got := minCostClimbingStairs([]int{10, 15, 20})
	want := 15

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMinCostClimbingStairs2(t *testing.T) {
	got := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	want := 6

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
