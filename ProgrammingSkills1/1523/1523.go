package main

// 1523. Count Odd Numbers in an Interval Range
func countOdds(low int, high int) int {
	var result int
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			result++
		}
	}
	return result
}
