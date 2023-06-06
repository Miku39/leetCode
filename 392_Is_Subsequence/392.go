package main

// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequence/?envType=study-plan&id=level-1
func isSubsequence(s string, t string) bool {
	index := 0
	for i := 0; i < len(s); i++ {
		isExist := false
		for j := index; j < len(t); j++ {
			if s[i] == t[j] {
				index = j + 1
				isExist = true
				break
			}
		}
		if !isExist {
			return false
		}
	}

	return true
}
