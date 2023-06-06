package main

// 733. Flood Fill
// https://leetcode.com/problems/flood-fill/?envType=study-plan&id=level-1
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	startColor := image[sr][sc]
	colorImage(image, sr, sc, color, startColor)
	return image
}

func colorImage(image [][]int, sr int, sc int, color int, searchColor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) {
		return
	}

	if image[sr][sc] != searchColor || image[sr][sc] == color {
		return
	}

	image[sr][sc] = color
	colorImage(image, sr-1, sc, color, searchColor)
	colorImage(image, sr+1, sc, color, searchColor)
	colorImage(image, sr, sc-1, color, searchColor)
	colorImage(image, sr, sc+1, color, searchColor)
}

/**
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	startColor := image[sr][sc]
	colorImage(startColor, image, sr, sc, color)
	return image
}

func colorImage(searchColor int, image [][]int, sr int, sc int, color int) {
	if image[sr][sc] == color {
		return
	}
	image[sr][sc] = color
	if sr-1 >= 0 && image[sr-1][sc] == searchColor {
		colorImage(searchColor, image, sr-1, sc, color)
	}
	if sr+1 < len(image) && image[sr+1][sc] == searchColor {
		colorImage(searchColor, image, sr+1, sc, color)
	}
	if sc-1 >= 0 && image[sr][sc-1] == searchColor {
		colorImage(searchColor, image, sr, sc-1, color)
	}
	if sc+1 < len(image[0]) && image[sr][sc+1] == searchColor {
		colorImage(searchColor, image, sr, sc+1, color)
	}
}
**/
