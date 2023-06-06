package main

import "testing"

func TestSortByBits(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	got := sortByBits(input)
	want := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSortByBits2(t *testing.T) {
	input := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	got := sortByBits(input)
	want := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSortByBits3(t *testing.T) {
	input := []int{10, 100, 1000, 10000}
	got := sortByBits(input)
	want := []int{10, 100, 10000, 1000}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
