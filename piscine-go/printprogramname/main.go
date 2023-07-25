package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	filePath := os.Args[0]
	for _, char := range filePath {
		if char != '.' && char != '/' {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune(rune('\n'))
}
