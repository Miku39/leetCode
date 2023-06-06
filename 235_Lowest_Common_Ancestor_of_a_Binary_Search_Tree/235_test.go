package main

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	p := &TreeNode{
		Val: 2,
	}
	q := &TreeNode{
		Val: 8,
	}
	got := lowestCommonAncestor(root, p, q)
	want := &TreeNode{
		Val: 6,
	}

	if got.Val != want.Val {
		t.Errorf("got %v, wanted %v", got.Val, want.Val)
	}
}
