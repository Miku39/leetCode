package main

// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/?envType=study-plan&id=level-1
func fib(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}

	return a + b
}
