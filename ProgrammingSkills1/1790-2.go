package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(areAlmostEqual2("npv", "zpn"))
	// fmt.Println(areAlmostEqual2("npv", "zpn"))
	// fmt.Println(areAlmostEqual2(
	// 	"eyapgosevqecyuobikhcpztfrrnuibtpjekygfssnqihtmlnevrwqosktoozmefq",
	// 	"iexryyjgmazbxegpxxipzoybexebydjwxgarqiysoqexgvelhlbpotlfbnogbvpg"),
	// )
}

// 1790. Check if One String Swap Can Make Strings Equal
func areAlmostEqual2(s1 string, s2 string) bool {
	array1 := strings.Split(s1, "")
	array2 := strings.Split(s2, "")

	diffCount := 0
	swapIndex1 := -1
	swapIndex2 := -1
	for i, v := range array1 {
		if v == array2[i] {
			continue
		}

		diffCount++

		if swapIndex1 < 0 {
			swapIndex1 = i
		} else {
			swapIndex2 = i
		}
	}

	if diffCount == 0 {
		return true
	}

	if diffCount != 2 {
		return false
	}

	if array1[swapIndex1] == array2[swapIndex2] && array1[swapIndex2] == array2[swapIndex1] {
		return true
	}

	return false
}
