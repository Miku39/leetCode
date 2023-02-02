package main

import "testing"

func TestGetDecimalValue(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	got := getDecimalValue(input)
	want := 5

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetDecimalValue2(t *testing.T) {
	input := &ListNode{
		Val:  0,
		Next: nil,
	}
	got := getDecimalValue(input)
	want := 0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
