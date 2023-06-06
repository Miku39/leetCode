package main

import (
	"strings"
)

// 205. Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/
func isIsomorphic(s string, t string) bool {
	stringS := strings.Split(s, "")
	stringT := strings.Split(t, "")

	stringMap := make(map[string]string)
	boolMap := make(map[string]bool)
	for i, sValue := range stringS {
		tValue, ok := stringMap[sValue]
		if !ok {
			tValue = stringT[i]
			isExist := boolMap[tValue]
			if isExist {
				return false
			}
			stringMap[sValue] = tValue
			boolMap[tValue] = true
			continue
		}

		if tValue != stringT[i] {
			return false
		}
	}

	return true
}
