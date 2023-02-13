package main

// 142. Linked List Cycle II
// https://leetcode.com/problems/linked-list-cycle-ii/?envType=study-plan&id=level-1

type ListNode struct {
	Val  int
	Next *ListNode
}

// @see https://leetcode.com/problems/linked-list-cycle-ii/solutions/44781/concise-o-n-solution-by-using-c-with-detailed-alogrithm-description/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head
	entry := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for {
				if entry == slow {
					return entry
				}
				entry = entry.Next
				slow = slow.Next
			}
		}
	}

	return nil
}
