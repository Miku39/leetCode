package main

import "testing"

// Input: root = [1,null,3,2,4,null,5,6]
// Output: [1,3,5,6,2,4]
func TestPreorder(t *testing.T) {
	input := &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val:      5,
						Children: []*Node{},
					},
					&Node{
						Val:      6,
						Children: []*Node{},
					},
				},
			},
			&Node{
				Val:      2,
				Children: []*Node{},
			},
			&Node{
				Val:      4,
				Children: []*Node{},
			},
		},
	}

	got := preorder(input)
	want := []int{1, 3, 5, 6, 2, 4}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Input: root = []
// Output: []
func TestPreorder2(t *testing.T) {
	got := preorder(nil)
	want := []int{}

	if !isSame(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func isSame(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
