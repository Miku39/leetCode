package main

import (
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	got := characterReplacement("ABAB", 2)
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCharacterReplacement2(t *testing.T) {
	got := characterReplacement("AABABBA", 1)
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCharacterReplacement3(t *testing.T) {
	got := characterReplacement("AAAA", 2)
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
