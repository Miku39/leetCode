package main

import "strings"

// 1768. Merge Strings Alternately
func mergeAlternately(word1 string, word2 string) string {
	array1 := strings.Split(word1, "")
	array2 := strings.Split(word2, "")

	length := len(array1)
	if length < len(array2) {
		length = len(array2)
	}

	var rersult string
	for i := 0; i < length; i++ {
		if i < len(array1) {
			rersult += array1[i]
		}
		if i < len(array2) {
			rersult += array2[i]
		}
	}

	return rersult
}
