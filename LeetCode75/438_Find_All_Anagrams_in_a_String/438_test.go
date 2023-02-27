package main

import (
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	got := findAnagrams("cbaebabacd", "abc")
	want := []int{0, 6}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFindAnagrams2(t *testing.T) {
	got := findAnagrams("abab", "ab")
	want := []int{0, 1, 2}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
