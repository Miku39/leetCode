package main

import "testing"

func TestMiddleNode(t *testing.T) {
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
	got := middleNode(input)
	want := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMiddleNode2(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	got := middleNode(input)
	want := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
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
