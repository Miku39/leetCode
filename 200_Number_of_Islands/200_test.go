package main

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	input := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	got := numIslands(input)
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestNumIslands2(t *testing.T) {
	input := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	got := numIslands(input)
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
