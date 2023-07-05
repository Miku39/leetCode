package main

import "testing"

func TestFindDifference(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	got := findDifference(nums1, nums2)
	want := [][]int{{1, 3}, {4, 6}}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFindDifference2(t *testing.T) {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	got := findDifference(nums1, nums2)
	want := [][]int{{3}, {}}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j, _ := range v {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
