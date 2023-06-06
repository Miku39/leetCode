package main

import (
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	got := topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2)
	want := []string{"i", "love"}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
