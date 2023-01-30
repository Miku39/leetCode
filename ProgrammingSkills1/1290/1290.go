package main

// 1290. Convert Binary Number in a Linked List to Integer
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result = result*2 + head.Val
		head = head.Next
	}

	return result
}
