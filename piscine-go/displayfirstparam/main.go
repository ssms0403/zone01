package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lastParam := os.Args[1]
	for _, c := range lastParam {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
