package main

// 404. Sum of Left Leaves
// https://leetcode.com/problems/sum-of-left-leaves/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return sumNumber(root, nil, 0)
}

func sumNumber(node *TreeNode, parentNode *TreeNode, number int) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil && node == parentNode.Left {
		number += node.Val
		return number
	}

	return sumNumber(node.Left, node, number) + sumNumber(node.Right, node, number)
}
