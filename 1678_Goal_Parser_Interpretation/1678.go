package main

import "strings"

// 1678. Goal Parser Interpretation
func interpret(command string) string {
	array := strings.Split(command, "")

	result := ""
	tmp := ""
	for _, v := range array {
		tmp += v
		switch tmp {
		case "G":
			result += "G"
			tmp = ""
		case "()":
			result += "o"
			tmp = ""
		case "(al)":
			result += "al"
			tmp = ""
		}
	}

	return result
}
