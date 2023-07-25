package piscine

func QuadC(
	x int,
	y int,
) {
	cornerTopRight := 'A'
	cornerTopLeft := cornerTopRight
	cornerBottomRight := 'C'
	cornerBottomLeft := cornerBottomRight
	edgeLine := 'B'
	edgeColumn := edgeLine
	drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}
