package main

// 589. N-ary Tree Preorder Traversal
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	result = appendChildrenVal(result, root.Children)
	return result
}

func appendChildrenVal(result []int, children []*Node) []int {
	for _, v := range children {
		result = append(result, v.Val)
		if len(v.Children) > 0 {
			result = appendChildrenVal(result, v.Children)
		}
	}
	return result
}
