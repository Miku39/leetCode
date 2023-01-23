package main

import "fmt"

func main() {
	result := countOdds(3, 7)
	fmt.Println(result)
}

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
