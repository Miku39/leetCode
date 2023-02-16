package main

import (
	"testing"
)

func TestLevelOrder(t *testing.T) {
	input := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	got := levelOrder(input)
	want := [][]int{{3}, {9, 20}, {15, 7}}

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
