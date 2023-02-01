package main

// 104. Maximum Depth of Binary Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return returnMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func returnMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
