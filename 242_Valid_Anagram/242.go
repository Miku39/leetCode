package main

import (
	"sort"
	"strings"
)

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	arrayS := strings.Split(s, "")
	arrayT := strings.Split(t, "")

	sort.Strings(arrayS)
	sort.Strings(arrayT)

	stringS := strings.Join(arrayS, "")
	stringT := strings.Join(arrayT, "")

	return stringS == stringT
}
