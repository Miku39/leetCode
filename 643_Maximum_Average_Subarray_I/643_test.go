package main

import "testing"

// func TestFindMaxAverage(t *testing.T) {
// 	nums := []int{1, 12, -5, -6, 50, 3}
// 	got := findMaxAverage(nums, 4)
// 	want := 12.75000

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}
// }

func TestFindMaxAverage2(t *testing.T) {
	nums := []int{4, 2, 1, 3, 3}
	got := findMaxAverage(nums, 2)
	want := 3.00000

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
