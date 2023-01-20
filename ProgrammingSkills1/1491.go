package main

import "fmt"

func main() {
	input := []int{4000, 3000, 1000, 2000}
	result := average(input)
	fmt.Println(result)
}

// 1491. Average Salary Excluding the Minimum and Maximum Salary
func average(salary []int) float64 {
	min := salary[0]
	max := salary[0]
	for _, num := range salary {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}

	sum := 0
	count := 0
	for _, num := range salary {
		if num == min || num == max {
			continue
		}
		sum += num
		count++
	}

	ave := float64(sum) / float64(count)
	return ave
}
