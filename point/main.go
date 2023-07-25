package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	str := "x = "
	printStr(str)
	str = ", y = "
	getNumber(points.x)
	printStr(str)
	getNumber(points.y)
	printStr("\n")
}

func getNumber(r int) {
	if r == 42 {
		z01.PrintRune(52)
		z01.PrintRune(50)
	} else {
		z01.PrintRune(50)
		z01.PrintRune(49)
	}
}
