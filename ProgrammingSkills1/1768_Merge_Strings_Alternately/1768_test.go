package main

import "testing"

func TestMergeAlternately(t *testing.T) {
	got := mergeAlternately("abc", "pqr")
	want := "apbqcr"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMergeAlternately2(t *testing.T) {
	got := mergeAlternately("ab", "pqrs")
	want := "apbqrs"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMergeAlternately3(t *testing.T) {
	got := mergeAlternately("abcd", "pq")
	want := "apbqcd"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
