package main

import "testing"

func TestNextGreaterElement(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	got := nextGreaterElement(nums1, nums2)
	want := []int{-1, 3, -1}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestNextGreaterElement2(t *testing.T) {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	got := nextGreaterElement(nums1, nums2)
	want := []int{3, -1}

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
