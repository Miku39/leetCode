package main

// 1232. Check If It Is a Straight Line
func checkStraightLine(coordinates [][]int) bool {

	baseX := coordinates[0][0]
	baseY := coordinates[0][1]
	x := coordinates[1][0]
	y := coordinates[1][1]

	diffX := x - baseX
	diffY := y - baseY

	for i := 1; i < len(coordinates); i++ {
		x = coordinates[i][0]
		y = coordinates[i][1]

		if (x-baseX)*diffY != (y-baseY)*diffX {
			return false
		}
	}

	return true
}
