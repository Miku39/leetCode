package main

import "testing"

func TestIsAlienSorted(t *testing.T) {
	got := isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsAlienSorted2(t *testing.T) {
	got := isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz")
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsAlienSorted3(t *testing.T) {
	got := isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz")
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
