package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(areAlmostEqual("abcd", "dcba"))
	fmt.Println(areAlmostEqual(
		"ysmpagrkzsmmzmsssutzgpxrmoylkgemgfcperptsxjcsgojwourhxlhqkxumonfgrczmjvbhwvhpnocz",
		"ysmpagrqzsmmzmsssutzgpxrmoylkgemgfcperptsxjcsgojwourhxlhkkxumonfgrczmjvbhwvhpnocz"),
	)
}

// 1790. Check if One String Swap Can Make Strings Equal
func areAlmostEqual(s1 string, s2 string) bool {
	array1 := strings.Split(s1, "")
	array2 := strings.Split(s2, "")

	swapCount := 0
	for i, v := range array1 {
		if v == array2[i] {
			continue
		}

		if swapCount != 0 {
			return false
		}
		swapCount++

		flag := false
		for j := i; j < len(array2); j++ {
			if v == array2[j] && array1[j] == array2[i] {
				array2[j] = array2[i]
				array2[i] = v
				flag = true
				break
			}
		}
		if flag == false {
			return false
		}
	}

	return true
}
