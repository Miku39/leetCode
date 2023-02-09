package main

import "testing"

func TestIsIsomorphic(t *testing.T) {
	got := isIsomorphic("egg", "add")
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsIsomorphic2(t *testing.T) {
	got := isIsomorphic("foo", "bar")
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsIsomorphic3(t *testing.T) {
	got := isIsomorphic("badc", "baba")
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
