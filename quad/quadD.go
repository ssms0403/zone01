package piscine

func QuadD(
	x int,
	y int,
) {
	cornerTopRight := 'C'
	cornerTopLeft := 'A'
	cornerBottomRight := cornerTopRight
	cornerBottomLeft := cornerTopLeft
	edgeLine := 'B'
	edgeColumn := edgeLine
	drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}
