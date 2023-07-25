package piscine

func QuadA(
	x int,
	y int,
) {
	cornerTopRight := 'o'
	cornerTopLeft := cornerTopRight
	cornerBottomRight := cornerTopRight
	cornerBottomLeft := cornerTopRight
	edgeLine := '-'
	edgeColumn := '|'
	drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}
