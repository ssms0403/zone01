package main

import "fmt"

type point struct {
	x int
	y int
}

type rectangle struct {
	upLeft    point
	downRight point
}

func defineRectangle(vPoint []point, n int) *rectangle {
	xMin := vPoint[0].x
	xMax := vPoint[0].x
	yMin := vPoint[0].y
	yMax := vPoint[0].y

	ptr := &rectangle{}

	for i := 0; i < n; i++ {
		if vPoint[i].x < xMin {
			xMin = vPoint[i].x
		}
		if vPoint[i].x > xMax {
			xMax = vPoint[i].x
		}
		if vPoint[i].y < yMin {
			yMin = vPoint[i].y
		}
		if vPoint[i].y > yMax {
			yMax = vPoint[i].y
		}
	}
	ptr.upLeft.x = xMin
	ptr.upLeft.y = yMax
	ptr.downRight.x = xMax
	ptr.downRight.y = yMin

	return ptr
}

func calcArea(ptr *rectangle) int {
	return (ptr.upLeft.x - ptr.downRight.x) * (ptr.downRight.y - ptr.upLeft.y)
}

func main() {
	vPoint := []point{}
	n := 7

	for i := 0; i < n; i++ {
		val := point{
			x: i%2 + 1,
			y: i + 2,
		}
		vPoint = append(vPoint, val)
	}

	rectangle := defineRectangle(vPoint, n)
	fmt.Println("area of the rectangle:", calcArea(rectangle))
}
