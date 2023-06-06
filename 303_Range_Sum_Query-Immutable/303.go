package main

// 303. Range Sum Query - Immutable
// https://leetcode.com/problems/range-sum-query-immutable/

type NumArray struct {
	sumArray []int
}

func Constructor(nums []int) NumArray {
	sumArray := make([]int, len(nums))
	sum := 0
	for i, num := range nums {
		sum += num
		sumArray[i] = sum
	}

	return NumArray{
		sumArray: sumArray,
	}
}

func (n *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return n.sumArray[right]
	}
	return n.sumArray[right] - n.sumArray[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
