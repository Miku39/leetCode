package main

import "testing"

func TestHammingWeight(t *testing.T) {
	got := hammingWeight(00000000000000000000000000001011)
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
