package main

// 70. Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i < n; i++ {
		a, b = b, a+b
	}

	return a + b
}
