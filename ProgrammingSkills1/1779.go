package main

import "fmt"

func main() {
	points := [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}
	result := nearestValidPoint(3, 4, points)
	fmt.Println(result)
}

// 1779. Find Nearest Point That Has the Same X or Y Coordinate
func nearestValidPoint(x int, y int, points [][]int) int {
	smallestDistance := 10000
	result := -1
	for i, point := range points {
		if x != point[0] && y != point[1] {
			continue
		}

		distance := abs(x-point[0]) + abs(y-point[1])
		if smallestDistance > distance {
			smallestDistance = distance
			result = i
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
