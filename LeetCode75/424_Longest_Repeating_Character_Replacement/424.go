package main

import (
	"strings"
)

// 424. Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/
func characterReplacement(s string, k int) int {
	sString := strings.Split(s, "")
	sMap := make(map[string]int)

	result := 0
	mostFreqLetter := 0
	left := 0
	for right, currentString := range sString {
		_, ok := sMap[currentString]
		if !ok {
			sMap[currentString] = 0
		}

		sMap[currentString]++

		if mostFreqLetter < sMap[currentString] {
			mostFreqLetter = sMap[currentString]
		}

		if (right-left+1)-mostFreqLetter > k {
			sMap[sString[left]] -= 1
			left += 1
		}

		if result < right-left+1 {
			result = right - left + 1
		}
	}

	return result
}
