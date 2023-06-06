package main

// 200. Number of Islands
// https://leetcode.com/problems/number-of-islands/?envType=study-plan&id=level-1
func numIslands(grid [][]byte) int {
	result := 0
	for row, v := range grid {
		for col := range v {
			if grid[row][col] == '1' {
				result++
				checkIsland(grid, row, col)
			}
		}
	}

	return result
}

func checkIsland(grid [][]byte, row int, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return
	}

	if grid[row][col] == '0' {
		return
	}

	grid[row][col] = '0'
	checkIsland(grid, row+1, col)
	checkIsland(grid, row-1, col)
	checkIsland(grid, row, col+1)
	checkIsland(grid, row, col-1)
}
