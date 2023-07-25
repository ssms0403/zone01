package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}
	args := os.Args[1:]

	if len(args)%2 == 0 {
		printStr(args[len(args)/2-1] + " " + args[len(args)/2])
	} else {
		printStr(args[len(args)/2])
	}
}

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
