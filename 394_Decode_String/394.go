package main

import (
	"strconv"
	"strings"
)

// 394. Decode String
// https://leetcode.com/problems/decode-string/
func decodeString(s string) string {
	arrayS := strings.Split(s, "")

	stackNumber := []int{}
	stackString := []string{}
	curNumber := 0
	curString := ""
	for i := 0; i < len(arrayS); i++ {
		if arrayS[i] == "[" {
			stackNumber = append(stackNumber, curNumber)
			stackString = append(stackString, curString)
			curNumber = 0
			curString = ""
		} else if arrayS[i] == "]" {
			repeatNumber := stackNumber[len(stackNumber)-1]
			stackNumber = stackNumber[:len(stackNumber)-1]
			prevString := stackString[len(stackString)-1]
			stackString = stackString[:len(stackString)-1]
			tmpString := prevString
			for j := 0; j < repeatNumber; j++ {
				tmpString += curString
			}
			curString = tmpString
		} else {
			number, err := strconv.Atoi(arrayS[i])
			if err != nil {
				curString += arrayS[i]
			} else {
				curNumber = curNumber*10 + number
			}
		}
	}

	return curString
}
