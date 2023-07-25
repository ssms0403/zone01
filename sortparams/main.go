package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for i := range arguments {
		for j := range arguments {
			if arguments[i] < arguments[j] {
				arguments[i], arguments[j] = arguments[j], arguments[i]
			}
		}
	}
	for _, arg := range arguments {
		for _, a := range arg {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
