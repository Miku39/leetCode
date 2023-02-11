package main

import "testing"

func TestMergeTwoLists(t *testing.T) {
	input1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	input2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	got := mergeTwoLists(input1, input2)
	want := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
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
