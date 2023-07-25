package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		for _, c := range arg {
			if c >= 'a' && c <= 'z' {
				z01.PrintRune(rune(c) - 32)
			} else if c >= 'A' && c <= 'Z' {
				z01.PrintRune(rune(c) + 32)
			} else {
				z01.PrintRune(rune(c))
			}
		}
		z01.PrintRune('\n')
	}
}
