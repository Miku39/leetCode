package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	got := containsDuplicate([]int{1, 2, 3, 1})
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
