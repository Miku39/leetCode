package main

import "testing"

func TestSubtractProductAndSum(t *testing.T) {
	got := subtractProductAndSum(4421)
	want := 21

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
