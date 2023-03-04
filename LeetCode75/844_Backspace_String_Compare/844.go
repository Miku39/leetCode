package main

import (
	"strings"
)

// 844. Backspace String Compare
// https://leetcode.com/problems/backspace-string-compare/
func backspaceCompare(s string, t string) bool {
	resultS := makeResult(s)
	resultT := makeResult(t)

	return resultS == resultT
}

func makeResult(input string) string {
	arrayInput := strings.Split(input, "")
	arrayResult := make([]string, 0, len(arrayInput))
	count := 0
	for i := len(arrayInput) - 1; i >= 0; i-- {
		if arrayInput[i] == "#" {
			count++
			continue
		}
		if count > 0 {
			count--
			continue
		}
		arrayResult = append(arrayResult, arrayInput[i])
	}

	return strings.Join(arrayResult, "")
}
