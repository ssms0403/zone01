package piscine

func QuadB(
	x int,
	y int,
) {
	cornerTopRight := '\\'
	cornerTopLeft := '/'
	cornerBottomRight := cornerTopLeft
	cornerBottomLeft := cornerTopRight
	edgeLine := '*'
	edgeColumn := edgeLine
	drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}
