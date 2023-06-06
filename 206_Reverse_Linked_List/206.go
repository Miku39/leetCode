package main

// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/?envType=study-plan&id=level-1

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var current *ListNode
	for head != nil {
		temp := &ListNode{
			Val:  head.Val,
			Next: current,
		}
		current = temp
		head = head.Next
	}
	return current
}

// func reverseList(head *ListNode) *ListNode {
// 	return helper(head, nil)
// }

// func helper(current *ListNode, prev *ListNode) *ListNode {
// 	if current == nil {
// 		return prev
// 	}
// 	next := current.Next
// 	current.Next = prev
// 	return helper(next, current)
// }
