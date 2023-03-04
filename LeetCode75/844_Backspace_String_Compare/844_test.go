package main

import (
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	got := backspaceCompare("ab#c", "ad#c")
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestBackspaceCompare2(t *testing.T) {
	got := backspaceCompare("ab##", "c#d#")
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestBackspaceCompare3(t *testing.T) {
	got := backspaceCompare("a#c", "b")
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
