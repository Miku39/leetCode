package main

import "testing"

func TestReverseVowels(t *testing.T) {
	input := "hello"
	got := reverseVowels(input)
	want := "holle"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestReverseVowels2(t *testing.T) {
	input := "Ui"
	got := reverseVowels(input)
	want := "iU"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
