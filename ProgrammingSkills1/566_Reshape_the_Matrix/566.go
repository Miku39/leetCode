package main

// 566. Reshape the Matrix
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	result := make([][]int, r)
	result[0] = make([]int, c)

	rowIndex := 0
	colIndex := 0
	for _, v := range mat {
		for _, w := range v {
			if colIndex == c {
				rowIndex++
				result[rowIndex] = make([]int, c)
				colIndex = 0
			}
			result[rowIndex][colIndex] = w
			colIndex++
		}
	}

	return result
}
