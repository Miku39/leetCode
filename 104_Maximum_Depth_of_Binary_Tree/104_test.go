package main

import "testing"

func TestMaxDepth(t *testing.T) {
	input := &TreeNode{
		Val: 1,
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
	got := maxDepth(input)
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
