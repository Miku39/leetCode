package main

// 1572. Matrix Diagonal Sum
func diagonalSum(mat [][]int) int {
	sum := 0
	for i, v := range mat {
		sum += v[i]
		sum += v[len(mat)-1-i]
	}

	if len(mat)%2 == 1 {
		midleIndex := len(mat) / 2
		sum -= mat[midleIndex][midleIndex]
	}

	return sum
}
