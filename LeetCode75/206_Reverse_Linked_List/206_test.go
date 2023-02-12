package main

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	got := reverseList(input)
	want := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMergeTwoLists2(t *testing.T) {
	got := reverseList(nil)

	if got != nil {
		t.Errorf("got %v, wanted %v", got, nil)
	}
}

func isSame(a, b *ListNode) bool {
	for {
		if a == nil && b == nil {
			return true
		}

		if (a != nil && b == nil) || (a == nil && b != nil) {
			return false
		}

		if a.Val != b.Val {
			return false
		}

		a = a.Next
		b = b.Next
	}
}
