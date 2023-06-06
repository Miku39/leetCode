package main

import "math"

// 98. Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/?envType=study-plan&id=level-1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt, math.MaxInt)
}

func validate(root *TreeNode, low int, high int) bool {
	if root == nil {
		return true
	}

	if low >= root.Val || root.Val >= high {
		return false
	}

	return validate(root.Left, low, root.Val) && validate(root.Right, root.Val, high)
}
