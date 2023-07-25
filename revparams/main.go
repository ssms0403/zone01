package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := len(arguments) - 1; i > 0; i-- {
		if i != 0 {
			for j := range arguments[i] {
				runes := []rune(arguments[i])
				z01.PrintRune(runes[j])
			}
			z01.PrintRune('\n')
		}
	}
}
