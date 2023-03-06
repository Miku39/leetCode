package main

import (
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	got := lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
