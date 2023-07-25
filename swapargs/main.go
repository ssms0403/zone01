package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 1 {
		for i := len(os.Args) - 1; i >= 1; i-- {
			printStr(os.Args[i])
			if i != 1 {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
