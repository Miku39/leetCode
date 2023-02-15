package main

import (
	"strings"
)

// 409. Longest Palindrome
// https://leetcode.com/problems/longest-palindrome/?envType=study-plan&id=level-1
func longestPalindrome(s string) int {
	stringArray := strings.Split(s, "")

	stringMap := make(map[string]int)
	for _, v := range stringArray {
		stringMap[v]++
	}

	result := 0
	oddCount := 0
	for _, w := range stringMap {
		if w%2 == 0 {
			result += w
		} else {
			result += w - 1
			oddCount++
		}
	}

	if oddCount > 0 {
		result++
	}

	return result
}
