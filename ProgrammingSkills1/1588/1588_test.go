package main

import "testing"

func TestSumOddLengthSubarrays(t *testing.T) {
	input := []int{1, 4, 2, 5, 3}
	got := sumOddLengthSubarrays(input)
	want := 58

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
