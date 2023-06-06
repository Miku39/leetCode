package main

import "testing"

func TestSumOfLeftLeaves(t *testing.T) {
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
	got := sumOfLeftLeaves(input)
	want := 24

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
