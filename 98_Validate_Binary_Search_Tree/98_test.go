package main

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	input := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	got := isValidBST(input)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsValidBST2(t *testing.T) {
	input := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	got := isValidBST(input)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
