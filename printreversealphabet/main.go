package main

import "github.com/01-edu/z01"

func main() {
	for i := 26; i > 0; i-- {
		z01.PrintRune(rune(96 + i))
	}
	z01.PrintRune('\n')
}
