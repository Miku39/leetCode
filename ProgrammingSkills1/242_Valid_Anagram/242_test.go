package main

import "testing"

func TestIsAnagram(t *testing.T) {
	got := isAnagram("anagram", "nagaram")
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsAnagram2(t *testing.T) {
	got := isAnagram("rat", "car")
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsAnagram3(t *testing.T) {
	got := isAnagram("a", "ab")
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
