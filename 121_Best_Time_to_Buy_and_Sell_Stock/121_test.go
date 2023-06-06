package main

import "testing"

func TestMaxProfit(t *testing.T) {
	got := maxProfit([]int{7, 1, 5, 3, 6, 4})
	want := 5

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMaxProfit2(t *testing.T) {
	got := maxProfit([]int{7, 6, 4, 3, 1})
	want := 0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
