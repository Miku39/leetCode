package main

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	got := decodeString("3[a]2[bc]")
	want := "aaabcbc"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDecodeString2(t *testing.T) {
	got := decodeString("3[a2[c]]")
	want := "accaccacc"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDecodeString3(t *testing.T) {
	got := decodeString("2[abc]3[cd]ef")
	want := "abcabccdcdcdef"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
