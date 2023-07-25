package piscine

func QuadE(
	x int,
	y int,
) {
	cornerTopRight := 'C'
	cornerTopLeft := 'A'
	cornerBottomRight := cornerTopLeft
	cornerBottomLeft := cornerTopRight
	edgeLine := 'B'
	edgeColumn := edgeLine
	drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}
