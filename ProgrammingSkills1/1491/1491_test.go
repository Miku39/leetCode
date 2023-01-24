package main

import "testing"

func TestAverage(t *testing.T) {
	got := average([]int{4000, 3000, 1000, 2000})
	want := 2500.00000

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
