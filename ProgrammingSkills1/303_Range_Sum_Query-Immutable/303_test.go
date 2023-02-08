package main

import "testing"

/*
*
Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Explanation
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
*
*/
func TestNumArray(t *testing.T) {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	got1 := numArray.SumRange(0, 2)
	want1 := 1
	if got1 != want1 {
		t.Errorf("got %v, wanted %v", got1, want1)
	}

	got2 := numArray.SumRange(2, 5)
	want2 := -1
	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

	got3 := numArray.SumRange(0, 5)
	want3 := -3
	if got3 != want3 {
		t.Errorf("got %v, wanted %v", got3, want3)
	}
}
