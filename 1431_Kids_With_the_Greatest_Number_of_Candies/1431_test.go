package main

import "testing"

func TestKidsWithCandies(t *testing.T) {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	got := kidsWithCandies(candies, extraCandies)
	want := []bool{true, true, true, false, true}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b []bool) bool {
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
