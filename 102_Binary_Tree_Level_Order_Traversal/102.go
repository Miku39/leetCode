package main

// 102. Binary Tree Level Order Traversal
// https://leetcode.com/problems/binary-tree-level-order-traversal/?envType=study-plan&id=level-1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	appendResult(&result, root, 0)
	return result
}

func appendResult(result *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(*result) <= level {
		*result = append(*result, []int{})
	}

	(*result)[level] = append((*result)[level], root.Val)

	level++
	appendResult(result, root.Left, level)
	appendResult(result, root.Right, level)
}
