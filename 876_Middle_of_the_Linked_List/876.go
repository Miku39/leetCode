package main

// 876. Middle of the Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	count := 0
	listNode := head
	for listNode != nil {
		count++
		listNode = listNode.Next
	}

	middleCount := count / 2

	result := head
	for i := 0; i < middleCount; i++ {
		result = result.Next
	}

	return result
}
