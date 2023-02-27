package main

import (
	"sort"
	"strings"
)

// 438. Find All Anagrams in a String
// https://leetcode.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	arrayS := strings.Split(s, "")
	arrayP := strings.Split(p, "")
	sort.Strings(arrayP)

	result := []int{}
	checkArray := make([]string, len(arrayP))
	for i := range arrayS {
		if len(arrayS)-len(arrayP) < i {
			break
		}
		for j := range arrayP {
			checkArray[j] = arrayS[i+j]
		}
		sort.Strings(checkArray)
		if strings.Join(arrayP, "") == strings.Join(checkArray, "") {
			result = append(result, i)
		}
	}
	return result
}
