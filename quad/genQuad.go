package piscine

import "github.com/01-edu/z01"

func drawQuad(
	maxColumn int,
	maxLine int,
	cornerTopRight rune,
	cornerTopLeft rune,
	cornerBottomRight rune,
	cornerBottomLeft rune,
	edgeLine rune,
	edgeColumn rune,
) {
	if maxLine <= 0 || maxColumn <= 0 {
		return
	}
	whiteSpace := rune(' ')
	for line := 1; line <= maxLine; line++ {
		if line == 1 {
			drawRow(cornerTopLeft, cornerTopRight, edgeLine, maxColumn)
		} else if line == maxLine {
			drawRow(cornerBottomLeft, cornerBottomRight, edgeLine, maxColumn)
		} else {
			drawRow(edgeColumn, edgeColumn, whiteSpace, maxColumn)
		}
	}
}

func drawRow(
	start rune,
	end rune,
	middle rune,
	maxColumn int,
) {
	for column := 1; column <= maxColumn; column++ {
		if column == 1 {
			z01.PrintRune(start)
		} else if column == maxColumn {
			z01.PrintRune(end)
		} else {
			z01.PrintRune(middle)
		}
	}
	z01.PrintRune('\n')
}
